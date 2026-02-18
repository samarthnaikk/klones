package genre

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateGenreRequest) (*Genre, error) {
	return s.repo.Create(req)
}

func (s *Service) List() ([]*Genre, error) {
	return s.repo.List()
}

func (s *Service) GetByID(id string) (*Genre, error) {
	return s.repo.GetByID(id)
}
