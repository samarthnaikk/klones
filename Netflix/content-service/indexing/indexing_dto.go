package indexing

type IndexContentRequest struct {
	ContentType string `json:"content_type"`
	Title       string `json:"title"`
	Tags        string `json:"tags"`
	Regions     string `json:"regions"`
}
