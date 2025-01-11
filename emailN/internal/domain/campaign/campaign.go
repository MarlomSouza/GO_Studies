package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"errors"
	"time"

	"github.com/rs/xid"
)

const (
	Pending string = "Pending"
	Started string = "Started"
	Done    string = "Done"
	Cancel  string = "Cancel"
)

type Contact struct {
	Id         string `gorm:"size:50"`
	Email      string `validate:"email" gorm:"size:100"`
	CampaignId string `gorm:"size:50`
}

type Campaign struct {
	Id         string    `validate:"required" gorm:"size:50" `
	Name       string    `validate:"min=5,max=24" gorm:"size:50"`
	CreatedOn  time.Time `validate:"required"`
	Content    string    `validate:"min=5,max=1024" gorm:"size:1050"`
	Recipients []Contact `validate:"required,dive"`
	Status     string    `gorm:"size:20"`
}

func NewCampaign(name string, content string, recipients []Contact) (*Campaign, error) {

	// isValid, err := isValidCampaign(name, content, recipients)
	// if !isValid {
	// 	return nil, err
	// }

	campaign := &Campaign{
		Id:         xid.New().String(),
		Name:       name,
		Content:    content,
		Recipients: recipients,
		CreatedOn:  time.Now(),
		Status:     Pending,
	}
	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
}

func (c *Campaign) Cancel() error {
	if c.Status != Pending {
		return errors.New("campaign status is invalid")
	}

	c.Status = Cancel
	return nil
}

// func isValidCampaign(name string, content string, recipients []Contact) (bool, error) {
// 	if name == "" {
// 		return false, errors.New("name is required")
// 	}

// 	if content == "" {
// 		return false, errors.New("content is required")
// 	}

// 	if len(recipients) == 0 {
// 		return false, errors.New("recipients is required")
// 	}

// 	return true, nil
// }

func NewContact(email string) (*Contact, error) {

	contact := &Contact{
		Email: email,
		Id:    xid.New().String(),
	}

	err := internalerrors.ValidateStruct(contact)

	if err == nil {
		return contact, nil
	}

	return nil, err
}
