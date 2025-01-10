package campaign

import (
	"emailn/internal/domain/campaign/contract"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Create_campaign(t *testing.T) {
	assert := assert.New(t)
	service := Service{}
	newCampaign := contract.NewCampaignDto{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"xxx@gmail.com", "xxx@outlook.com"},
	}

	campaignId, err := service.Create(newCampaign)

	assert.NotNil(campaignId)
	assert.Nil(err)
}
