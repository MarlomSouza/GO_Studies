package campaign

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)
 
var (
	name	= "New Function"
	content	= "Use new function in go"
	emails 	= []Contact{{ Email: "xxx@gmail.com"}}
)

func Test_NewCampaign(t *testing.T) {
	assert:= assert.New(t)

	sut := NewCampaign(name, content, emails)

	assert.Equal(name, sut.Name)
	assert.Equal(content, sut.Content)
	assert.Equal(emails, sut.Contacts)
}

func Test_NewCampaign_IdIsNotNil(t *testing.T) {
	assert:= assert.New(t)

	sut := NewCampaign(name, content, emails)

	assert.NotNil(sut.Id)
}

func Test_NewCampaign_CreatedOnIsNotNil(t *testing.T) {
	assert:= assert.New(t)

	sut := NewCampaign(name, content, emails)
	
	assert.WithinDuration(time.Now(), sut.CreatedOn, 5 * time.Second)
}