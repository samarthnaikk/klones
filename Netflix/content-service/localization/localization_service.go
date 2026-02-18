package localization

type Service struct {
	repo *Repository
}

func NewService(repo *Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetByContentID(contentID, lang string) ([]*Localization, error) {
	return s.repo.GetByContentID(contentID, lang)
}

func (s *Service) Upsert(contentID string, req UpdateLocalizationRequest) (*Localization, error) {
	return s.repo.Upsert(contentID, req)
}
