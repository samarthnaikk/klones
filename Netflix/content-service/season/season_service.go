package season

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateSeasonRequest) (*Season, error) {
	return s.repo.Create(req)
}

func (s *Service) GetByID(id string) (*Season, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListBySeriesID(seriesID string) ([]*Season, error) {
	return s.repo.ListBySeriesID(seriesID)
}

func (s *Service) Update(id string, req UpdateSeasonRequest) (*Season, error) {
	return s.repo.Update(id, req)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
