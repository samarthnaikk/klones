package episode

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

func (r *Repository) Create(req CreateEpisodeRequest) (*Episode, error) {
	e := &Episode{
		ID:              newID(),
		SeasonID:        req.SeasonID,
		EpisodeNumber:   req.EpisodeNumber,
		Title:           req.Title,
		Description:     req.Description,
		DurationMinutes: req.DurationMinutes,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
	}
	_, err := r.db.Exec(
		`INSERT INTO episodes (id, season_id, episode_number, title, description, duration_minutes, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)`,
		e.ID, e.SeasonID, e.EpisodeNumber, e.Title, e.Description, e.DurationMinutes, e.CreatedAt, e.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Repository) GetByID(id string) (*Episode, error) {
	e := &Episode{}
	err := r.db.QueryRow(`SELECT id, season_id, episode_number, title, description, duration_minutes, created_at, updated_at FROM episodes WHERE id=$1`, id).
		Scan(&e.ID, &e.SeasonID, &e.EpisodeNumber, &e.Title, &e.Description, &e.DurationMinutes, &e.CreatedAt, &e.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Repository) ListBySeasonID(seasonID string) ([]*Episode, error) {
	rows, err := r.db.Query(`SELECT id, season_id, episode_number, title, description, duration_minutes, created_at, updated_at FROM episodes WHERE season_id=$1`, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Episode
	for rows.Next() {
		e := &Episode{}
		if err := rows.Scan(&e.ID, &e.SeasonID, &e.EpisodeNumber, &e.Title, &e.Description, &e.DurationMinutes, &e.CreatedAt, &e.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, e)
	}
	return result, nil
}

func (r *Repository) Update(id string, req UpdateEpisodeRequest) (*Episode, error) {
	e, err := r.GetByID(id)
	if err != nil {
		return nil, err
	}
	if req.EpisodeNumber != nil {
		e.EpisodeNumber = *req.EpisodeNumber
	}
	if req.Title != nil {
		e.Title = *req.Title
	}
	if req.Description != nil {
		e.Description = *req.Description
	}
	if req.DurationMinutes != nil {
		e.DurationMinutes = *req.DurationMinutes
	}
	e.UpdatedAt = time.Now()
	_, err = r.db.Exec(
		`UPDATE episodes SET episode_number=$1, title=$2, description=$3, duration_minutes=$4, updated_at=$5 WHERE id=$6`,
		e.EpisodeNumber, e.Title, e.Description, e.DurationMinutes, e.UpdatedAt, e.ID,
	)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func (r *Repository) Delete(id string) error {
	_, err := r.db.Exec(`DELETE FROM episodes WHERE id=$1`, id)
	return err
}
