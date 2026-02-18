package metadata

type UpdateMetadataRequest struct {
	ContentType       *string `json:"content_type,omitempty"`
	Tags              *string `json:"tags,omitempty"`
	Language          *string `json:"language,omitempty"`
	SubtitleLanguages *string `json:"subtitle_languages,omitempty"`
	AudioLanguages    *string `json:"audio_languages,omitempty"`
}
