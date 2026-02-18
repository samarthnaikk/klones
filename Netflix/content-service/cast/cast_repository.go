package cast

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

func (r *Repository) Create(req CreateCastRequest) (*Cast, error) {
	c := &Cast{
		ID:          newID(),
		ContentID:   req.ContentID,
		ContentType: req.ContentType,
		PersonName:  req.PersonName,
		Role:        req.Role,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO cast_members (id, content_id, content_type, person_name, role, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7)`,
		c.ID, c.ContentID, c.ContentType, c.PersonName, c.Role, c.CreatedAt, c.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (r *Repository) ListByContentID(contentID string) ([]*Cast, error) {
	rows, err := r.db.Query(`SELECT id, content_id, content_type, person_name, role, created_at, updated_at FROM cast_members WHERE content_id=$1`, contentID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Cast
	for rows.Next() {
		c := &Cast{}
		if err := rows.Scan(&c.ID, &c.ContentID, &c.ContentType, &c.PersonName, &c.Role, &c.CreatedAt, &c.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}
