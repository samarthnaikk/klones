package availability

import "time"

type UpdateAvailabilityRequest struct {
	Region         *string    `json:"region,omitempty"`
	AvailableFrom  *time.Time `json:"available_from,omitempty"`
	AvailableUntil *time.Time `json:"available_until,omitempty"`
	LicenseType    *string    `json:"license_type,omitempty"`
}
