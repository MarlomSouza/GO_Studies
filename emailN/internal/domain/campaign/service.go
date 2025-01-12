package campaign

import (
	"emailn/internal/contract"

	internalerrors "emailn/internal/internal-errors"
)

type Service interface {
	Create(dto contract.NewCampaignDto) (string, error)
	Get() ([]contract.CampaignDto, error)
	GetById(id string) (*contract.CampaignDto, error)
	Cancel(id string) error
	Delete(id string) error
}

type ServiceImp struct {
	Repository Repository
}

// Create creates a new campaign and saves it to the repository.
func (s *ServiceImp) Create(dto contract.NewCampaignDto) (string, error) {
	recipients := []Contact{}

	for _, email := range dto.Emails {
		recipient, _ := NewContact(email)
		recipients = append(recipients, *recipient)
	}

	campaign, err := NewCampaign(dto.Name, dto.Content, recipients, dto.CreatedBy)
	if err != nil {
		return "", err
	}

	err = s.Repository.Create(campaign)
	if err != nil {
		return "", internalerrors.ErrInternal
	}

	return campaign.Id, nil
}

// Get retrieves all campaigns from the repository and converts them to DTOs.
func (s *ServiceImp) Get() ([]contract.CampaignDto, error) {
	campaigns, err := s.Repository.Get()

	if err != nil {
		return []contract.CampaignDto{}, internalerrors.ErrInternal
	}

	campaignDtos := make([]contract.CampaignDto, 0, len(campaigns))
	for _, campaign := range campaigns {
		recipientEmails := make([]string, 0, len(campaign.Recipients))
		for _, recipient := range campaign.Recipients {
			recipientEmails = append(recipientEmails, recipient.Email)
		}
		campaignDtos = append(campaignDtos, contract.CampaignDto{
			Name:    campaign.Name,
			Content: campaign.Content,
			Emails:  recipientEmails,
			Status:  campaign.Status,
		})
	}

	return campaignDtos, nil
}

// GetById retrieves a campaign by its ID from the repository and converts it to a DTO.
func (s *ServiceImp) GetById(id string) (*contract.CampaignDto, error) {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return nil, internalerrors.ProcessInternalError(err)
	}

	recipientEmails := make([]string, 0, len(campaign.Recipients))
	for _, recipient := range campaign.Recipients {
		recipientEmails = append(recipientEmails, recipient.Email)
	}

	campaignDtos := contract.CampaignDto{
		Name:    campaign.Name,
		Content: campaign.Content,
		Emails:  recipientEmails,
		Status:  campaign.Status,
	}

	return &campaignDtos, nil
}

func (s *ServiceImp) Cancel(id string) error {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return internalerrors.ProcessInternalError(err)
	}

	err = campaign.Cancel()

	if err != nil {
		return err
	}

	err = s.Repository.Update(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}
	return nil

}

func (s *ServiceImp) Delete(id string) error {
	campaign, err := s.Repository.GetById(id)

	if err != nil {
		return internalerrors.ProcessInternalError(err)
	}

	err = s.Repository.Delete(campaign)

	if err != nil {
		return internalerrors.ErrInternal
	}

	return nil
}
