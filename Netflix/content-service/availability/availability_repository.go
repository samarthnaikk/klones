package availability

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

func newID() string {
	return uuid.New().String()
}

func (r *Repository) GetByContentID(contentID, region string) ([]*Availability, error) {
	query := `SELECT id, content_id, region, available_from, available_until, license_type, created_at, updated_at FROM availability WHERE content_id=$1`
	args := []interface{}{contentID}
	if region != "" {
		query += " AND region=$2"
		args = append(args, region)
	}
	rows, err := r.db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var result []*Availability
	for rows.Next() {
		a := &Availability{}
		if err := rows.Scan(&a.ID, &a.ContentID, &a.Region, &a.AvailableFrom, &a.AvailableUntil, &a.LicenseType, &a.CreatedAt, &a.UpdatedAt); err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil
}

func (r *Repository) Upsert(contentID string, req UpdateAvailabilityRequest) (*Availability, error) {
	a := &Availability{
		ID:        newID(),
		ContentID: contentID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	if req.Region != nil {
		a.Region = *req.Region
	}
	if req.AvailableFrom != nil {
		a.AvailableFrom = *req.AvailableFrom
	}
	if req.AvailableUntil != nil {
		a.AvailableUntil = *req.AvailableUntil
	}
	if req.LicenseType != nil {
		a.LicenseType = *req.LicenseType
	}
	_, err := r.db.Exec(
		`INSERT INTO availability (id, content_id, region, available_from, available_until, license_type, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		 ON CONFLICT (content_id, region) DO UPDATE SET available_from=$4, available_until=$5, license_type=$6, updated_at=$8`,
		a.ID, a.ContentID, a.Region, a.AvailableFrom, a.AvailableUntil, a.LicenseType, a.CreatedAt, a.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}
	return a, nil
}
