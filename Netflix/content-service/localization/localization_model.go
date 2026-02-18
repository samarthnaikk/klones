package localization

import "time"

type Localization struct {
	ID          string    `json:"id" db:"id"`
	ContentID   string    `json:"content_id" db:"content_id"`
	Language    string    `json:"language" db:"language"`
	Title       string    `json:"title" db:"title"`
	Description string    `json:"description" db:"description"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
