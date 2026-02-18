package movie

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

func (r *Repository) Create(req CreateMovieRequest) (*Movie, error) {
	m := &Movie{
		ID:              newID(),
		Title:           req.Title,
		Description:     req.Description,
		ReleaseYear:     req.ReleaseYear,
		DurationMinutes: req.DurationMinutes,
		GenreIDs:        req.GenreIDs,
		Rating:          req.Rating,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO movies (id, title, description, release_year, duration_minutes, genre_ids, rating, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
		m.ID, m.Title, m.Description, m.ReleaseYear, m.DurationMinutes, m.GenreIDs, m.Rating, m.CreatedAt, m.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) GetByID(id string) (*Movie, error) {
	m := &Movie{}
	err := r.db.QueryRow(`SELECT id, title, description, release_year, duration_minutes, genre_ids, rating, created_at, updated_at FROM movies WHERE id=$1`, id).
		Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseYear, &m.DurationMinutes, &m.GenreIDs, &m.Rating, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) List() ([]*Movie, error) {
	rows, err := r.db.Query(`SELECT id, title, description, release_year, duration_minutes, genre_ids, rating, created_at, updated_at FROM movies`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var movies []*Movie
	for rows.Next() {
		m := &Movie{}
		if err := rows.Scan(&m.ID, &m.Title, &m.Description, &m.ReleaseYear, &m.DurationMinutes, &m.GenreIDs, &m.Rating, &m.CreatedAt, &m.UpdatedAt); err != nil {
			return nil, err
		}
		movies = append(movies, m)
	}
	return movies, nil
}

func (r *Repository) Update(id string, req UpdateMovieRequest) (*Movie, error) {
	m, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.Title != nil {
		m.Title = *req.Title
	}
	if req.Description != nil {
		m.Description = *req.Description
	}
	if req.ReleaseYear != nil {
		m.ReleaseYear = *req.ReleaseYear
	}
	if req.DurationMinutes != nil {
		m.DurationMinutes = *req.DurationMinutes
	}
	if req.GenreIDs != nil {
		m.GenreIDs = *req.GenreIDs
	}
	if req.Rating != nil {
		m.Rating = *req.Rating
	}
	m.UpdatedAt = time.Now()
	_, err = r.db.Exec(
		`UPDATE movies SET title=$1, description=$2, release_year=$3, duration_minutes=$4, genre_ids=$5, rating=$6, updated_at=$7 WHERE id=$8`,
		m.Title, m.Description, m.ReleaseYear, m.DurationMinutes, m.GenreIDs, m.Rating, m.UpdatedAt, m.ID,
	)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM movies WHERE id=$1`, id)
	return err
}
