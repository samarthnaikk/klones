package series

type CreateSeriesRequest struct {
	Title        string  `json:"title"`
	Description  string  `json:"description"`
	ReleaseYear  int     `json:"release_year"`
	GenreIDs     string  `json:"genre_ids"`
	TotalSeasons int     `json:"total_seasons"`
	Rating       float64 `json:"rating"`
}

type UpdateSeriesRequest struct {
	Title        *string  `json:"title,omitempty"`
	Description  *string  `json:"description,omitempty"`
	ReleaseYear  *int     `json:"release_year,omitempty"`
	GenreIDs     *string  `json:"genre_ids,omitempty"`
	TotalSeasons *int     `json:"total_seasons,omitempty"`
	Rating       *float64 `json:"rating,omitempty"`
}
