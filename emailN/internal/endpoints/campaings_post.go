package endpoints

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/render"
)

func (h *HandlerCampaign) CampaignPost(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {

	var request contract.NewCampaignDto
	err := render.DecodeJSON(r.Body, &request)
	if err != nil {
		println(err.Error())
	}

	id, err := h.CampaignService.Create(request)

	return EndpointStruct{
		Obj: map[string]string{"id": id}, Status: 201,
	}, err

}
