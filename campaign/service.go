package campaign

type Service interface {
	GetCampaigns(userID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID == 0 {
		campaign, err := s.repository.FindAll()
		return campaign, err
	}

	campaign, err := s.repository.FindByUserID(userID)
	return campaign, err
}
