package cast

type CreateCastRequest struct {
	ContentID   string `json:"content_id"`
	ContentType string `json:"content_type"`
	PersonName  string `json:"person_name"`
	Role        string `json:"role"`
}
