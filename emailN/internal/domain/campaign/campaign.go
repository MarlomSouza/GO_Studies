package campaign

import "time"


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
		Id: "1",
		Name:  name,
		Content: content,
		Contacts:  emails,
		CreatedOn: time.Now(),
	}
}
