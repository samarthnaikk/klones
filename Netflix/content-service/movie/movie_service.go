package movie

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateMovieRequest) (*Movie, error) {
	return s.repo.Create(req)
}

func (s *Service) GetByID(id string) (*Movie, error) {
	return s.repo.GetByID(id)
}

func (s *Service) List() ([]*Movie, error) {
	return s.repo.List()
}

func (s *Service) Update(id string, req UpdateMovieRequest) (*Movie, error) {
	return s.repo.Update(id, req)
}

func (s *Service) Delete(id string) error {
	return s.repo.Delete(id)
}
