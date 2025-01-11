package endpoints

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func (h *HandlerCampaign) CampaignCancelPatch(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {

	id := chi.URLParam(r, "id")

	err := h.CampaignService.Cancel(id)

	return EndpointStruct{
		Obj: nil, Status: http.StatusOK,
	}, err

}
