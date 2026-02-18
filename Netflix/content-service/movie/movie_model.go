package movie

import "time"

type Movie struct {
	ID              string    `json:"id" db:"id"`
	Title           string    `json:"title" db:"title"`
	Description     string    `json:"description" db:"description"`
	ReleaseYear     int       `json:"release_year" db:"release_year"`
	DurationMinutes int       `json:"duration_minutes" db:"duration_minutes"`
	GenreIDs        string    `json:"genre_ids" db:"genre_ids"`
	Rating          float64   `json:"rating" db:"rating"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}
