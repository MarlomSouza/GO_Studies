package endpoints

import (
	"bytes"
	"context"
	"emailn/internal/contract"
	internalmock "emailn/internal/test/internal-mock"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/jaswdr/faker/v2"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var fake = faker.New()

func setup(body contract.NewCampaignDto, expectedCreatedBy string) (*httptest.ResponseRecorder, *http.Request) {

	var buf bytes.Buffer
	json.NewEncoder(&buf).Encode(body)

	req, _ := http.NewRequest(http.MethodPost, "/", &buf)
	ctx := context.WithValue(req.Context(), "email", expectedCreatedBy)
	req = req.WithContext(ctx)
	res := httptest.NewRecorder()
	return res, req

}

func Test_CampaignPost_should_save_new_campaign(t *testing.T) {
	assert := assert.New(t)
	body := contract.NewCampaignDto{
		Name:    "test create",
		Content: fake.Lorem().Text(100),
		Emails:  []string{"xxx@gmail.com"},
	}
	expectedCreatedBy := "xxx@gmail.com"
	service := new(internalmock.CampaignServiceMock)
	service.On("Create", mock.MatchedBy(func(request contract.NewCampaignDto) bool {
		return request.Name == body.Name && request.Content == body.Content && request.CreatedBy == expectedCreatedBy
	})).Return("1", nil)

	handler := HandlerCampaign{CampaignService: service}
	res, req := setup(body, expectedCreatedBy)

	response, _ := handler.CampaignPost(res, req)

	assert.Equal(http.StatusCreated, response.Status)

}
