package localization

type UpdateLocalizationRequest struct {
	Language    *string `json:"language,omitempty"`
	Title       *string `json:"title,omitempty"`
	Description *string `json:"description,omitempty"`
}
