package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
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

func (r *repositoryMock) Get() error {
	args := r.Called()
	return args.Error(0)

}

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"xxx@gmail.com", "xxx@outlook.com"},
	}
	service = Service{}
)

func Test_Create_campaign(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockRepository

	campaignId, err := service.Create(newCampaign)

	assert.NotNil(campaignId)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockRepository

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_SaveCampaign_ValidateObject(t *testing.T) {
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.MatchedBy(func(campaign *Campaign) bool {
		return newCampaign.Name == campaign.Name && newCampaign.Content == campaign.Content
	})).Return(nil)
	service.Repository = mockRepository

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_ValidateDomainError(t *testing.T) {
	assert := assert.New(t)

	_, err := service.Create(contract.NewCampaignDto{})

	assert.False(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Create_ValidateDatabaseError(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(repositoryMock)
	mockRepository.On("Save", mock.Anything).Return(errors.New("error while saving in database"))
	service.Repository = mockRepository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}
