package localization

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func newID() string {
	return uuid.New().String()
}

func (r *Repository) GetByContentID(contentID, lang string) ([]*Localization, error) {
	query := `SELECT id, content_id, language, title, description, created_at, updated_at FROM localizations WHERE content_id=$1`
	args := []interface{}{contentID}
	if lang != "" {
		query += " AND language=$2"
		args = append(args, lang)
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Localization
	for rows.Next() {
		l := &Localization{}
		if err := rows.Scan(&l.ID, &l.ContentID, &l.Language, &l.Title, &l.Description, &l.CreatedAt, &l.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, l)
	}
	return result, nil
}

func (r *Repository) Upsert(contentID string, req UpdateLocalizationRequest) (*Localization, error) {
	l := &Localization{
		ID:        newID(),
		ContentID: contentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if req.Language != nil {
		l.Language = *req.Language
	}
	if req.Title != nil {
		l.Title = *req.Title
	}
	if req.Description != nil {
		l.Description = *req.Description
	}
	_, err := r.db.Exec(
		`INSERT INTO localizations (id, content_id, language, title, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)
		 ON CONFLICT (content_id, language) DO UPDATE SET title=$4, description=$5, updated_at=$7`,
		l.ID, l.ContentID, l.Language, l.Title, l.Description, l.CreatedAt, l.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return l, nil
}
