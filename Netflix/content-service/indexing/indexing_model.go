package indexing

import "time"

type Index struct {
	ID            string    `json:"id" db:"id"`
	ContentID     string    `json:"content_id" db:"content_id"`
	ContentType   string    `json:"content_type" db:"content_type"`
	Title         string    `json:"title" db:"title"`
	Tags          string    `json:"tags" db:"tags"`
	Regions       string    `json:"regions" db:"regions"`
	Indexed       bool      `json:"indexed" db:"indexed"`
	LastIndexedAt time.Time `json:"last_indexed_at" db:"last_indexed_at"`
	CreatedAt     time.Time `json:"created_at" db:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" db:"updated_at"`
}
