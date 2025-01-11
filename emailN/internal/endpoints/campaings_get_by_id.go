package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerCampaign) CampaignGetById(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {
	id := chi.URLParam(r, "id")
	campaign, err := h.CampaignService.GetById(id)

	if campaign == nil && err == nil {
		return EndpointStruct{
			Obj: nil, Status: http.StatusNotFound,
		}, nil
	}

	return EndpointStruct{
		Obj: campaign, Status: http.StatusOK,
	}, err

}
