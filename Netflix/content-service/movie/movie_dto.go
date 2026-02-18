package movie

type CreateMovieRequest struct {
	Title           string  `json:"title"`
	Description     string  `json:"description"`
	ReleaseYear     int     `json:"release_year"`
	DurationMinutes int     `json:"duration_minutes"`
	GenreIDs        string  `json:"genre_ids"`
	Rating          float64 `json:"rating"`
}

type UpdateMovieRequest struct {
	Title           *string  `json:"title,omitempty"`
	Description     *string  `json:"description,omitempty"`
	ReleaseYear     *int     `json:"release_year,omitempty"`
	DurationMinutes *int     `json:"duration_minutes,omitempty"`
	GenreIDs        *string  `json:"genre_ids,omitempty"`
	Rating          *float64 `json:"rating,omitempty"`
}
