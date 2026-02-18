package series

import "time"

type Series struct {
	ID           string    `json:"id" db:"id"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	ReleaseYear  int       `json:"release_year" db:"release_year"`
	GenreIDs     string    `json:"genre_ids" db:"genre_ids"`
	TotalSeasons int       `json:"total_seasons" db:"total_seasons"`
	Rating       float64   `json:"rating" db:"rating"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
