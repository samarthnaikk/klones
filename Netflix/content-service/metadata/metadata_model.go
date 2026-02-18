package metadata

import "time"

type Metadata struct {
	ID                string    `json:"id" db:"id"`
	ContentID         string    `json:"content_id" db:"content_id"`
	ContentType       string    `json:"content_type" db:"content_type"`
	Tags              string    `json:"tags" db:"tags"`
	Language          string    `json:"language" db:"language"`
	SubtitleLanguages string    `json:"subtitle_languages" db:"subtitle_languages"`
	AudioLanguages    string    `json:"audio_languages" db:"audio_languages"`
	CreatedAt         time.Time `json:"created_at" db:"created_at"`
	UpdatedAt         time.Time `json:"updated_at" db:"updated_at"`
}
