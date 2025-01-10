package endpoints

import (
	"bytes"
	"emailn/internal/contract"
	"emailn/internal/domain/campaign"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var fake = faker.New()

type serviceMock struct {
	mock.Mock
}

func (r *serviceMock) Create(dto contract.NewCampaignDto) (string, error) {
	args := r.Called(dto)
	return args.String(0), args.Error(1)

}
func (r *serviceMock) Get() ([]campaign.Campaign, error) {
	args := r.Called()
	return args.Get(0).([]campaign.Campaign), args.Error(1)
}

func (r *serviceMock) GetById(id string) (contract.CampaignDto, error) {
	args := r.Called(id)
	return args.Get(0).(contract.CampaignDto), args.Error(1)
}

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignDto{
		Name:    "test create",
		Content: fake.Lorem().Text(100),
		Emails:  []string{"xxx@gmail.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignDto) bool {
		return request.Name == body.Name && request.Content == body.Content
	})).Return("1", nil)

	handler := HandlerCampaign{CampaignService: service}

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	res := httptest.NewRecorder()

	response, err := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, response.Status)
	assert.Nil(err)
}

func Test_CampaignPost_should_inform_error_when_exist(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignDto{
		Name:    "test create",
		Content: fake.Lorem().Text(100),
		Emails:  []string{"xxx@gmail.com"},
	}
	service := new(serviceMock)
	service.On("Create", mock.Anything).Return("", errors.New("Error when creating"))
	handler := HandlerCampaign{CampaignService: service}
	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)
	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	res := httptest.NewRecorder()

	_, sut := handler.CampaignPost(res, req)

	assert.NotNil(sut)
}
