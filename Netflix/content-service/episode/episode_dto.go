package episode

type CreateEpisodeRequest struct {
	SeasonID        string `json:"season_id"`
	EpisodeNumber   int    `json:"episode_number"`
	Title           string `json:"title"`
	Description     string `json:"description"`
	DurationMinutes int    `json:"duration_minutes"`
}

type UpdateEpisodeRequest struct {
	EpisodeNumber   *int    `json:"episode_number,omitempty"`
	Title           *string `json:"title,omitempty"`
	Description     *string `json:"description,omitempty"`
	DurationMinutes *int    `json:"duration_minutes,omitempty"`
}
