package campaign

import (
	"testing"

	"github.com/stretchr/testify/assert"
)
 
func TestNewCampaign(t *testing.T) {
	assert:= assert.New(t)
	name := "New Function"
	content := "Use new function in go"
	emails := []Contact{{
		Email: "xxx@gmail.com",
	}}
	
	sut := NewCampaign(name, content, emails)

	assert.Equal(name, sut.Name)
	assert.Equal(content, sut.Content)
	assert.Equal(emails, sut.Contacts)
}