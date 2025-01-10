package campaign

import (
	"emailn/internal/domain/campaign/contract"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (r *repositoryMock) Save(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)

}

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"xxx@gmail.com", "xxx@outlook.com"},
	}

	mockRepository = new(repositoryMock)
	service        = Service{repository: mockRepository}
)

func Test_Create_campaign(t *testing.T) {
	assert := assert.New(t)
	mockRepository.On("Save", mock.Anything).Return(nil)

	campaignId, err := service.Create(newCampaign)

	assert.NotNil(campaignId)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	mockRepository.On("Save", mock.Anything).Return(nil)

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_SaveCampaign_ValidateObject(t *testing.T) {
	mockRepository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return newCampaign.Name == campaign.Name && newCampaign.Content == campaign.Content
	})).Return(nil)

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)
	newCampaign.Name = ""

	_, err := service.Create(newCampaign)

	assert.Equal("name is required", err.Error())
}

func Test_Create_ValidateDatabaseError(t *testing.T) {
	assert := assert.New(t)
	mockRepository.On("Save", mock.Anything).Return(errors.New("error while saving in database"))

	_, err := service.Create(newCampaign)

	assert.Equal("error while saving in database", err.Error())
}
