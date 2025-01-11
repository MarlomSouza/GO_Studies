package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerCampaign) CampaignGetById(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {
	id := chi.URLParam(r, "id")
	campaigns, err := h.CampaignService.GetById(id)

	return EndpointStruct{
		Obj: campaigns, Status: 200,
	}, err

}
