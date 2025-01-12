package endpoints

import (
	internalerrors "emailn/internal/internal-errors"
	internalmock "emailn/internal/test/internal-mock"

	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func Test_Delete_Should_delete_campaign(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	service.On("Delete", mock.Anything).Return(nil)
	handler := HandlerCampaign{CampaignService: service}

	req, _ := http.NewRequest(http.MethodDelete, "/{id}", nil)
	res := httptest.NewRecorder()

	sut, err := handler.CampaignDelete(res, req)

	assert.Equal(http.StatusNoContent, sut.Status)
	assert.Nil(err)
}

func Test_Delete_Return_not_found_when_campaign_does_not_exist(t *testing.T) {
	assert := assert.New(t)
	service := new(internalmock.CampaignServiceMock)
	expectedError := internalerrors.ErrNotFound
	service.On("Delete", mock.Anything).Return(expectedError)
	handler := HandlerCampaign{CampaignService: service}
	req, _ := http.NewRequest(http.MethodDelete, "/{id}", nil)
	res := httptest.NewRecorder()

	_, err := handler.CampaignDelete(res, req)

	assert.Equal(expectedError.Error(), err.Error())

}
