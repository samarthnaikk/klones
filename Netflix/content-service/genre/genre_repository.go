package genre

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

func (r *Repository) Create(req CreateGenreRequest) (*Genre, error) {
	g := &Genre{
		ID:          newID(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO genres (id, name, description, created_at, updated_at) VALUES ($1,$2,$3,$4,$5)`,
		g.ID, g.Name, g.Description, g.CreatedAt, g.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return g, nil
}

func (r *Repository) List() ([]*Genre, error) {
	rows, err := r.db.Query(`SELECT id, name, description, created_at, updated_at FROM genres`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Genre
	for rows.Next() {
		g := &Genre{}
		if err := rows.Scan(&g.ID, &g.Name, &g.Description, &g.CreatedAt, &g.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, g)
	}
	return result, nil
}

func (r *Repository) GetByID(id string) (*Genre, error) {
	g := &Genre{}
	err := r.db.QueryRow(`SELECT id, name, description, created_at, updated_at FROM genres WHERE id=$1`, id).
		Scan(&g.ID, &g.Name, &g.Description, &g.CreatedAt, &g.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return g, nil
}
