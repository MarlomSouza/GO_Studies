package campaign

import (
	"emailn/internal/contract"
	internalerrors "emailn/internal/internal-errors"
	"errors"

	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	newCampaign = contract.NewCampaignDto{
		Name:    "Test Y",
		Content: "Content",
		Emails:  []string{"xxx@gmail.com", "xxx@outlook.com"},
	}
	service = ServiceImp{}
)

type CampaignRepositoryMock struct {
	mock.Mock
}

func (r *CampaignRepositoryMock) Create(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(0)

}

func (r *CampaignRepositoryMock) Get() ([]Campaign, error) {
	args := r.Called()
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]Campaign), nil

}

func (r *CampaignRepositoryMock) GetById(id string) (*Campaign, error) {
	args := r.Called(id)
	if args.Error(1) != nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Campaign), args.Error(1)
}

func (r *CampaignRepositoryMock) Update(campaign *Campaign) error {
	args := r.Called()

	return args.Error(1)
}

func (r *CampaignRepositoryMock) Delete(campaign *Campaign) error {
	args := r.Called(campaign)
	return args.Error(1)
}

func Test_Create_campaign(t *testing.T) {
	assert := assert.New(t)
	mockRepository := new(CampaignRepositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockRepository

	campaignId, err := service.Create(newCampaign)

	assert.NotNil(campaignId)
	assert.Nil(err)
}

func Test_Create_SaveCampaign(t *testing.T) {
	mockRepository := new(CampaignRepositoryMock)
	mockRepository.On("Save", mock.Anything).Return(nil)
	service.Repository = mockRepository

	service.Create(newCampaign)

	mockRepository.AssertExpectations(t)
}

func Test_Create_SaveCampaign_ValidateObject(t *testing.T) {
	mockRepository := new(CampaignRepositoryMock)
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
	mockRepository := new(CampaignRepositoryMock)
	mockRepository.On("Save", mock.Anything).Return(errors.New("error while saving in database"))
	service.Repository = mockRepository

	_, err := service.Create(newCampaign)

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Get_ShouldReturnAllCampaigns(t *testing.T) {
	assert := assert.New(t)
	expectedObject := []Campaign{
		{
			Name:       fake.Company().Name(),
			Content:    fake.Lorem().Text(100),
			Status:     Pending,
			Recipients: []Contact{{Email: "xxx@gmail.com"}},
		},
		{
			Name:       fake.Company().Name(),
			Content:    fake.Lorem().Text(100),
			Status:     Pending,
			Recipients: []Contact{{Email: "uuu@gmail.com"}},
		},
	}
	mockRepository := new(CampaignRepositoryMock)
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
	mockRepository := new(CampaignRepositoryMock)
	mockRepository.On("Get").Return([]Campaign{}, errors.New("not found"))
	service.Repository = mockRepository

	_, err := service.Get()

	assert.True(errors.Is(err, internalerrors.ErrInternal))
}

func Test_Get_ShouldReturnACampaignBasedOnId(t *testing.T) {
	assert := assert.New(t)
	expectedCampaign, _ := NewCampaign(fake.Company().Name(), fake.Lorem().Text(100), []Contact{{Email: "xxx@gmail.com"}})
	mockRepository := new(CampaignRepositoryMock)
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
	mockRepository := new(CampaignRepositoryMock)
	mockRepository.On("GetById", mock.MatchedBy(func(id string) bool {
		return expectedId == id
	})).Return(Campaign{}, errors.New("not found"))
	service.Repository = mockRepository

	_, sut := service.GetById(expectedId)

	assert.True(errors.Is(sut, internalerrors.ErrInternal))

}
