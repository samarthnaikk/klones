package season

type CreateSeasonRequest struct {
	SeriesID     string `json:"series_id"`
	SeasonNumber int    `json:"season_number"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	ReleaseYear  int    `json:"release_year"`
}

type UpdateSeasonRequest struct {
	SeasonNumber *int    `json:"season_number,omitempty"`
	Title        *string `json:"title,omitempty"`
	Description  *string `json:"description,omitempty"`
	ReleaseYear  *int    `json:"release_year,omitempty"`
}
