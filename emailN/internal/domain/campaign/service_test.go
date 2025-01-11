package campaign_test

import (
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	internalerrors "emailn/internal/internal-errors"
	"emailn/internal/test/internalmock"
	"errors"

	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"xxx@gmail.com", "xxx@outlook.com"},
	}
	service = campaign.ServiceImp{}
	fake    = faker.New()
)

func Test_Create_campaign(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Create", mock.Anything).Return(nil)
	service.Repository = mockRepository

	campaignId, err := service.Create(newCampaign)

	assert.NotNil(campaignId)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Create", mock.Anything).Return(nil)
	service.Repository = mockRepository

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_SaveCampaign_ValidateObject(t *testing.T) {
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Create", mock.MatchedBy(func(campaign *campaign.Campaign) bool {
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
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Create", mock.Anything).Return(errors.New("error while saving in database"))
	service.Repository = mockRepository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Get_ShouldReturnAllCampaigns(t *testing.T) {
	assert := assert.New(t)
	expectedObject := []campaign.Campaign{
		{
			Name:       fake.Company().Name(),
			Content:    fake.Lorem().Text(100),
			Status:     campaign.Pending,
			Recipients: []campaign.Contact{{Email: "xxx@gmail.com"}},
		},
		{
			Name:       fake.Company().Name(),
			Content:    fake.Lorem().Text(100),
			Status:     campaign.Pending,
			Recipients: []campaign.Contact{{Email: "uuu@gmail.com"}},
		},
	}
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Get").Return(expectedObject, nil)
	service.Repository = mockRepository

	sut, err := service.Get()

	assert.Equal(expectedObject[0].Name, sut[0].Name)
	assert.Equal(expectedObject[0].Content, sut[0].Content)
	assert.Equal(expectedObject[0].Status, sut[0].Status)
	assert.Equal(expectedObject[0].Recipients[0].Email, sut[0].Emails[0])
	assert.Nil(err)

}

func Test_Get_ShouldReturnNilWhenNoCampaigns(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("Get").Return([]campaign.Campaign{}, errors.New("not found"))
	service.Repository = mockRepository

	_, err := service.Get()

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Get_ShouldReturnACampaignBasedOnId(t *testing.T) {
	assert := assert.New(t)
	expectedCampaign, _ := campaign.NewCampaign(fake.Company().Name(), fake.Lorem().Text(100), []campaign.Contact{{Email: "xxx@gmail.com"}})
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return id == expectedCampaign.Id
	})).Return(expectedCampaign, nil)
	service.Repository = mockRepository

	sut, err := service.GetById(expectedCampaign.Id)

	assert.Equal(expectedCampaign.Name, sut.Name)
	assert.Equal(expectedCampaign.Content, sut.Content)
	assert.Equal(expectedCampaign.Status, sut.Status)
	assert.Equal(expectedCampaign.Recipients[0].Email, sut.Emails[0])
	assert.Nil(err)
}

func Test_Get_ShouldReturnNilWhenIdNotFound(t *testing.T) {
	assert := assert.New(t)
	expectedId := "2"
	mockRepository := new(internalmock.CampaignRepositoryMock)
	mockRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return expectedId == id
	})).Return(campaign.Campaign{}, errors.New("not found"))
	service.Repository = mockRepository

	_, sut := service.GetById(expectedId)

	assert.True(errors.Is(sut, internalerrors.ErrInternal))

}
