package season

import "time"

type Season struct {
	ID           string    `json:"id" db:"id"`
	SeriesID     string    `json:"series_id" db:"series_id"`
	SeasonNumber int       `json:"season_number" db:"season_number"`
	Title        string    `json:"title" db:"title"`
	Description  string    `json:"description" db:"description"`
	ReleaseYear  int       `json:"release_year" db:"release_year"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`
}
