package metadata

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

func (r *Repository) GetByContentID(contentID string) (*Metadata, error) {
	m := &Metadata{}
	err := r.db.QueryRow(`SELECT id, content_id, content_type, tags, language, subtitle_languages, audio_languages, created_at, updated_at FROM metadata WHERE content_id=$1`, contentID).
		Scan(&m.ID, &m.ContentID, &m.ContentType, &m.Tags, &m.Language, &m.SubtitleLanguages, &m.AudioLanguages, &m.CreatedAt, &m.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func (r *Repository) Upsert(contentID string, req UpdateMetadataRequest) (*Metadata, error) {
	existing, err := r.GetByContentID(contentID)
	if err != nil {
		m := &Metadata{
			ID:        newID(),
			ContentID: contentID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		if req.ContentType != nil {
			m.ContentType = *req.ContentType
		}
		if req.Tags != nil {
			m.Tags = *req.Tags
		}
		if req.Language != nil {
			m.Language = *req.Language
		}
		if req.SubtitleLanguages != nil {
			m.SubtitleLanguages = *req.SubtitleLanguages
		}
		if req.AudioLanguages != nil {
			m.AudioLanguages = *req.AudioLanguages
		}
		_, err = r.db.Exec(
			`INSERT INTO metadata (id, content_id, content_type, tags, language, subtitle_languages, audio_languages, created_at, updated_at) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9)`,
			m.ID, m.ContentID, m.ContentType, m.Tags, m.Language, m.SubtitleLanguages, m.AudioLanguages, m.CreatedAt, m.UpdatedAt,
		)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	if req.ContentType != nil {
		existing.ContentType = *req.ContentType
	}
	if req.Tags != nil {
		existing.Tags = *req.Tags
	}
	if req.Language != nil {
		existing.Language = *req.Language
	}
	if req.SubtitleLanguages != nil {
		existing.SubtitleLanguages = *req.SubtitleLanguages
	}
	if req.AudioLanguages != nil {
		existing.AudioLanguages = *req.AudioLanguages
	}
	existing.UpdatedAt = time.Now()
	_, err = r.db.Exec(
		`UPDATE metadata SET content_type=$1, tags=$2, language=$3, subtitle_languages=$4, audio_languages=$5, updated_at=$6 WHERE content_id=$7`,
		existing.ContentType, existing.Tags, existing.Language, existing.SubtitleLanguages, existing.AudioLanguages, existing.UpdatedAt, contentID,
	)
	if err != nil {
		return nil, err
	}
	return existing, nil
}
