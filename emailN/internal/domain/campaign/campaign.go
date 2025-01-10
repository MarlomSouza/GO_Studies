package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string `validate:"required"`
}

type Campaign struct {
	Id         string    `validate:"required"`
	Name       string    `validate:"min=5,max=24"`
	CreatedOn  time.Time `validate:"required"`
	Content    string    `validate:"min=5,max=24"`
	Recipients []Contact `validate:"required"`
}

func NewCampaign(name string, content string, recipients []Contact) (*Campaign, error) {

	// isValid, err := isValidCampaign(name, content, recipients)
	// if !isValid {
	// 	return nil, err
	// }

	return &Campaign{
		Id:         xid.New().String(),
		Name:       name,
		Content:    content,
		Recipients: recipients,
		CreatedOn:  time.Now(),
	}, nil
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

	if email == "" {
		return nil, errors.New("invalid email")
	}

	return &Contact{
		Email: email,
	}, nil
}
