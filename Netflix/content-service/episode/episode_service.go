package episode

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateEpisodeRequest) (*Episode, error) {
	return s.repo.Create(req)
}

func (s *Service) GetByID(id string) (*Episode, error) {
	return s.repo.GetByID(id)
}

func (s *Service) ListBySeasonID(seasonID string) ([]*Episode, error) {
	return s.repo.ListBySeasonID(seasonID)
}

func (s *Service) Update(id string, req UpdateEpisodeRequest) (*Episode, error) {
	return s.repo.Update(id, req)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
