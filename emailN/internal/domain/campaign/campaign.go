package campaign

import (
	internalerrors "emailn/internal/internal-errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"email"`
}

type Campaign struct {
	Id         string    `validate:"required"`
	Name       string    `validate:"min=5,max=24"`
	CreatedOn  time.Time `validate:"required"`
	Content    string    `validate:"min=5,max=1024"`
	Recipients []Contact `validate:"required,dive"`
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
	}
	err := internalerrors.ValidateStruct(campaign)

	if err == nil {
		return campaign, nil
	}

	return nil, err
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
	}

	err := internalerrors.ValidateStruct(contact)

	if err == nil {
		return contact, nil
	}

	return nil, err
}
