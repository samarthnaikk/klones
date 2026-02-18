package availability

import "time"

type Availability struct {
	ID             string    `json:"id" db:"id"`
	ContentID      string    `json:"content_id" db:"content_id"`
	Region         string    `json:"region" db:"region"`
	AvailableFrom  time.Time `json:"available_from" db:"available_from"`
	AvailableUntil time.Time `json:"available_until" db:"available_until"`
	LicenseType    string    `json:"license_type" db:"license_type"`
	CreatedAt      time.Time `json:"created_at" db:"created_at"`
	UpdatedAt      time.Time `json:"updated_at" db:"updated_at"`
}
