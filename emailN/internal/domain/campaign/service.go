package campaign

import (
	"emailn/internal/domain/campaign/contract"
)

type Service struct {
	repository Repository
}

func (s *Service) Create(dto contract.NewCampaignDto) (string, error) {
	recipients := []Contact{}

	for _, email := range dto.Emails {
		recipient, _ := NewContact(email)
		recipients = append(recipients, *recipient)
	}

	campaign, err := NewCampaign(dto.Name, dto.Content, recipients)
	if err != nil {
		return "", err
	}

	err = s.repository.Save(campaign)
	if err != nil {
		return "", err
	}

	return campaign.Id, err
}
