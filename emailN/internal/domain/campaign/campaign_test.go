package campaign

import (
	"testing"
)
 
func TestNewCampaign(t *testing.T) {
	name := "New Function"
	content := "Use new function in go "
	emails := []Contact{{
		Email: "xxx@gmail.com",
	}}
	
	sut := NewCampaign(name, content, emails)

	if sut.Id != "1"{
		t.Errorf("expected 1")
	}
	
	if(sut.Name != name) {
		t.Errorf("Name if diff")
	}

}