package metadata

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByContentID(contentID string) (*Metadata, error) {
	return s.repo.GetByContentID(contentID)
}

func (s *Service) Upsert(contentID string, req UpdateMetadataRequest) (*Metadata, error) {
	return s.repo.Upsert(contentID, req)
}
