package campaign

import (
	"testing"
	"time"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
)

var (
	name      = "New Function"
	createdBy = faker.Faker.Person(faker.New()).Contact().Email
	content   = "Use new function in go"
	emails    = []Contact{{Email: "xxx@gmail.com"}}
	fake      = faker.New()
)

func Test_NewCampaign(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(name, sut.Name)
	assert.Equal(content, sut.Content)
	assert.Equal(emails, sut.Recipients)
	assert.Equal(createdBy, sut.CreatedBy)
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails, createdBy)

	assert.NotNil(sut.Id)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert := assert.New(t)

	sut, _ := NewCampaign(name, content, emails, createdBy)

	assert.WithinDuration(time.Now(), sut.CreatedOn, 5*time.Second)
}

func Test_NewCampaign_MustValidateNameWithMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign("", content, emails, createdBy)

	assert.Equal("name is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateNameWithMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(fake.Lorem().Text(1025), content, emails, createdBy)

	assert.Equal("name is required with max 24", err.Error())
}

func Test_NewCampaign_MustValidateContentMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, "", emails, createdBy)

	assert.Equal("content is required with min 5", err.Error())
}

func Test_NewCampaign_MustValidateContentMax(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, fake.Lorem().Text(1060), emails, createdBy)

	assert.Equal("content is required with max 1024", err.Error())
}

func Test_NewCampaign_MustValidateRecipientsMin(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, nil, createdBy)

	assert.Equal("recipients is required", err.Error())
}

func Test_NewCampaign_MustValidateCreatedBy(t *testing.T) {
	assert := assert.New(t)

	_, err := NewCampaign(name, content, emails, "ss")

	assert.Equal("createdby is invalid", err.Error())
}

func Test_NewCampaign_MustStartWithStatusPending(t *testing.T) {
	assert := assert.New(t)

	campaign, _ := NewCampaign(name, content, emails, createdBy)

	assert.Equal(Pending, campaign.Status)
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

	assert.Equal("email is invalid", err.Error())
}
