package indexing

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) List() ([]*Index, error) {
	return s.repo.List()
}

func (s *Service) IndexContent(contentID string, req IndexContentRequest) (*Index, error) {
	return s.repo.IndexContent(contentID, req)
}
