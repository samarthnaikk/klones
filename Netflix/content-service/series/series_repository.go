package series

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

func (r *Repository) Create(req CreateSeriesRequest) (*Series, error) {
	s := &Series{
		ID:           newID(),
		Title:        req.Title,
		Description:  req.Description,
		ReleaseYear:  req.ReleaseYear,
		GenreIDs:     req.GenreIDs,
		TotalSeasons: req.TotalSeasons,
		Rating:       req.Rating,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO series (id, title, description, release_year, genre_ids, total_seasons, rating, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		s.ID, s.Title, s.Description, s.ReleaseYear, s.GenreIDs, s.TotalSeasons, s.Rating, s.CreatedAt, s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) GetByID(id string) (*Series, error) {
	s := &Series{}
	err := r.db.QueryRow(`SELECT id, title, description, release_year, genre_ids, total_seasons, rating, created_at, updated_at FROM series WHERE id=$1`, id).
		Scan(&s.ID, &s.Title, &s.Description, &s.ReleaseYear, &s.GenreIDs, &s.TotalSeasons, &s.Rating, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) List() ([]*Series, error) {
	rows, err := r.db.Query(`SELECT id, title, description, release_year, genre_ids, total_seasons, rating, created_at, updated_at FROM series`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Series
	for rows.Next() {
		s := &Series{}
		if err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.ReleaseYear, &s.GenreIDs, &s.TotalSeasons, &s.Rating, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (r *Repository) Update(id string, req UpdateSeriesRequest) (*Series, error) {
	s, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		s.Title = *req.Title
	}
	if req.Description != nil {
		s.Description = *req.Description
	}
	if req.ReleaseYear != nil {
		s.ReleaseYear = *req.ReleaseYear
	}
	if req.GenreIDs != nil {
		s.GenreIDs = *req.GenreIDs
	}
	if req.TotalSeasons != nil {
		s.TotalSeasons = *req.TotalSeasons
	}
	if req.Rating != nil {
		s.Rating = *req.Rating
	}
	s.UpdatedAt = time.Now()
	_, err = r.db.Exec(
		`UPDATE series SET title=$1, description=$2, release_year=$3, genre_ids=$4, total_seasons=$5, rating=$6, updated_at=$7 WHERE id=$8`,
		s.Title, s.Description, s.ReleaseYear, s.GenreIDs, s.TotalSeasons, s.Rating, s.UpdatedAt, s.ID,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM series WHERE id=$1`, id)
	return err
}
