package season

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

func (r *Repository) Create(req CreateSeasonRequest) (*Season, error) {
	s := &Season{
		ID:           newID(),
		SeriesID:     req.SeriesID,
		SeasonNumber: req.SeasonNumber,
		Title:        req.Title,
		Description:  req.Description,
		ReleaseYear:  req.ReleaseYear,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO seasons (id, series_id, season_number, title, description, release_year, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		s.ID, s.SeriesID, s.SeasonNumber, s.Title, s.Description, s.ReleaseYear, s.CreatedAt, s.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) GetByID(id string) (*Season, error) {
	s := &Season{}
	err := r.db.QueryRow(`SELECT id, series_id, season_number, title, description, release_year, created_at, updated_at FROM seasons WHERE id=$1`, id).
		Scan(&s.ID, &s.SeriesID, &s.SeasonNumber, &s.Title, &s.Description, &s.ReleaseYear, &s.CreatedAt, &s.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) ListBySeriesID(seriesID string) ([]*Season, error) {
	rows, err := r.db.Query(`SELECT id, series_id, season_number, title, description, release_year, created_at, updated_at FROM seasons WHERE series_id=$1`, seriesID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Season
	for rows.Next() {
		s := &Season{}
		if err := rows.Scan(&s.ID, &s.SeriesID, &s.SeasonNumber, &s.Title, &s.Description, &s.ReleaseYear, &s.CreatedAt, &s.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, s)
	}
	return result, nil
}

func (r *Repository) Update(id string, req UpdateSeasonRequest) (*Season, error) {
	s, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.SeasonNumber != nil {
		s.SeasonNumber = *req.SeasonNumber
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
	s.UpdatedAt = time.Now()
	_, err = r.db.Exec(
		`UPDATE seasons SET season_number=$1, title=$2, description=$3, release_year=$4, updated_at=$5 WHERE id=$6`,
		s.SeasonNumber, s.Title, s.Description, s.ReleaseYear, s.UpdatedAt, s.ID,
	)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM seasons WHERE id=$1`, id)
	return err
}
