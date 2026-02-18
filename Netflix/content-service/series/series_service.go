package series

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateSeriesRequest) (*Series, error) {
	return s.repo.Create(req)
}

func (s *Service) GetByID(id string) (*Series, error) {
	return s.repo.GetByID(id)
}

func (s *Service) List() ([]*Series, error) {
	return s.repo.List()
}

func (s *Service) Update(id string, req UpdateSeriesRequest) (*Series, error) {
	return s.repo.Update(id, req)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
