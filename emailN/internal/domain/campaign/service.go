package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
)

type Service interface {
	Create(dto contract.NewCampaignDto) (string, error)
	Get() ([]Campaign, error)
	GetById(id string) (contract.CampaignDto, error)
}

type ServiceImp struct {
	Repository Repository
}

func (s *ServiceImp) Create(dto contract.NewCampaignDto) (string, error) {
	recipients := []Contact{}

	for _, email := range dto.Emails {
		recipient, _ := NewContact(email)
		recipients = append(recipients, *recipient)
	}

	campaign, err := NewCampaign(dto.Name, dto.Content, recipients)
	if err != nil {
		return "", err
	}

	err = s.Repository.Save(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.Id, nil
}

func (s *ServiceImp) Get() ([]Campaign, error) {
	campaigns, err := s.Repository.Get()

	if err != nil {
		return []Campaign{}, internalerrors.ErrInternal
	}

	return campaigns, nil
}

func (s *ServiceImp) GetById(id string) (contract.CampaignDto, error) {
	campaigns, err := s.Repository.GetById(id)

	if err != nil {
		return contract.CampaignDto{}, internalerrors.ErrInternal
	}

	dto := contract.CampaignDto{
		Name:    campaigns.Name,
		Content: campaigns.Content,
		Status:  campaigns.Status,
	}

	return dto, nil
}
