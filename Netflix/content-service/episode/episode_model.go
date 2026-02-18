package episode

import "time"

type Episode struct {
	ID              string    `json:"id" db:"id"`
	SeasonID        string    `json:"season_id" db:"season_id"`
	EpisodeNumber   int       `json:"episode_number" db:"episode_number"`
	Title           string    `json:"title" db:"title"`
	Description     string    `json:"description" db:"description"`
	DurationMinutes int       `json:"duration_minutes" db:"duration_minutes"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
