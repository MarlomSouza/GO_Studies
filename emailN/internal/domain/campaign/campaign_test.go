package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	name    = "New Function"
	content = "Use new function in go"
	emails  = []Contact{{Email: "xxx@gmail.com"}}
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails)

	assert.Equal(name, sut.Name)
	assert.Equal(content, sut.Content)
	assert.Equal(emails, sut.Recipients)
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails)

	assert.NotNil(sut.Id)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails)

	assert.WithinDuration(time.Now(), sut.CreatedOn, 5*time.Second)
}

func Test_NewCampaign_MustValidateName(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails)

	assert.Equal("name is required", err.Error())
}

func Test_NewCampaign_MustValidateContent(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails)

	assert.Equal("content is required", err.Error())
}

func Test_NewCampaign_MustValidateRecipients(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, []Contact{})

	assert.Equal("recipients is required", err.Error())
}

func Test_NewContact(t *testing.T) {
	assert := assert.New(t)
	email := "xxx@gmail.com"

	sut, _ := NewContact(email)

	assert.Equal(email, sut.Email)
}

func Test_NewContact_MustNotBeEmpty(t *testing.T) {
	assert := assert.New(t)

	_, err := NewContact("")

	assert.Equal("invalid email", err.Error())
}
