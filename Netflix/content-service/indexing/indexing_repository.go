package indexing

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

func (r *Repository) List() ([]*Index, error) {
	rows, err := r.db.Query(`SELECT id, content_id, content_type, title, tags, regions, indexed, last_indexed_at, created_at, updated_at FROM content_index`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Index
	for rows.Next() {
		idx := &Index{}
		if err := rows.Scan(&idx.ID, &idx.ContentID, &idx.ContentType, &idx.Title, &idx.Tags, &idx.Regions, &idx.Indexed, &idx.LastIndexedAt, &idx.CreatedAt, &idx.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, idx)
	}
	return result, nil
}

func (r *Repository) IndexContent(contentID string, req IndexContentRequest) (*Index, error) {
	idx := &Index{
		ID:            newID(),
		ContentID:     contentID,
		ContentType:   req.ContentType,
		Title:         req.Title,
		Tags:          req.Tags,
		Regions:       req.Regions,
		Indexed:       true,
		LastIndexedAt: time.Now(),
		CreatedAt:     time.Now(),
		UpdatedAt:     time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO content_index (id, content_id, content_type, title, tags, regions, indexed, last_indexed_at, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10)
		 ON CONFLICT (content_id) DO UPDATE SET content_type=$3, title=$4, tags=$5, regions=$6, indexed=$7, last_indexed_at=$8, updated_at=$10`,
		idx.ID, idx.ContentID, idx.ContentType, idx.Title, idx.Tags, idx.Regions, idx.Indexed, idx.LastIndexedAt, idx.CreatedAt, idx.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return idx, nil
}
