package endpoints

import (
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"

	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Delete_Should_return_campaign(t *testing.T) {
	assert := assert.New(t)
	campaignDto := contract.CampaignDto{
		Name:    "test create",
		Content: fake.Lorem().Text(100),
		Emails:  []string{"cxxx@gmail.com"},
	}
	service := new(internalmock.CampaignServiceMock)
	service.On("GetById", mock.Anything).Return(&campaignDto, nil)
	handler := HandlerCampaign{CampaignService: service}

	req, _ := http.NewRequest(http.MethodGet, "/{id}", nil)
	res := httptest.NewRecorder()

	sut, err := handler.CampaignGetById(res, req)

	assert.Equal(http.StatusOK, sut.Status)
	assert.Equal(campaignDto.Name, sut.Obj.(*contract.CampaignDto).Name)
	assert.Equal(campaignDto.Content, sut.Obj.(*contract.CampaignDto).Content)
	assert.Nil(err)
}

func Test_Delete_Should_return_error_when_something_wrong(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	expectedError := errors.New("Error when fetching")
	service.On("GetById", mock.Anything).Return(nil, expectedError)
	handler := HandlerCampaign{CampaignService: service}
	req, _ := http.NewRequest(http.MethodGet, "/{id}", nil)
	res := httptest.NewRecorder()

	_, sut := handler.CampaignGetById(res, req)

	assert.Equal(expectedError.Error(), sut.Error())
}
