package cast

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(req CreateCastRequest) (*Cast, error) {
	return s.repo.Create(req)
}

func (s *Service) ListByContentID(contentID string) ([]*Cast, error) {
	return s.repo.ListByContentID(contentID)
}
