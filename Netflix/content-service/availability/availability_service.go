package availability

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByContentID(contentID, region string) ([]*Availability, error) {
	return s.repo.GetByContentID(contentID, region)
}

func (s *Service) Upsert(contentID string, req UpdateAvailabilityRequest) (*Availability, error) {
	return s.repo.Upsert(contentID, req)
}
