package cast

import "time"

type Cast struct {
	ID          string    `json:"id" db:"id"`
	ContentID   string    `json:"content_id" db:"content_id"`
	ContentType string    `json:"content_type" db:"content_type"`
	PersonName  string    `json:"person_name" db:"person_name"`
	Role        string    `json:"role" db:"role"`
	CreatedAt   time.Time `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time `json:"updated_at" db:"updated_at"`
}
