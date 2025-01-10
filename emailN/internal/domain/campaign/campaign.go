package campaign

import (
	"errors"
	"time"

	"github.com/rs/xid"
)

type Contact struct {
	Email string
}

type Campaign struct {
	Id        string
	Name      string
	CreatedOn time.Time
	Content   string
	Recipients []Contact

}

func NewCampaign (name string, content string, recipients []Contact) (*Campaign, error){

	isValid, err := isValidCampaign(name, content, recipients)
	if !isValid {
		return nil, err
	}

	return &Campaign{
		Id: xid.New().String(),
		Name:  name,
		Content: content,
		Recipients:  recipients,
		CreatedOn: time.Now(),
	}, nil
}


func isValidCampaign(name string, content string, recipients []Contact) (bool, error) {
	if name == "" {
		return false, errors.New("name is required")
	}

	if content == "" {
		return false, errors.New("content is required")
	}

	if len(recipients) == 0 {
		return false, errors.New("recipients is required")
	}

	return true,nil
}
