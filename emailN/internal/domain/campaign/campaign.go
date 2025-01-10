package campaign

import (
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
	Contacts []Contact

}

func NewCampaign (name string, content string, emails []Contact) *Campaign{
	return &Campaign{
		Id: xid.New().String(),
		Name:  name,
		Content: content,
		Contacts:  emails,
		CreatedOn: time.Now(),
	}
}
