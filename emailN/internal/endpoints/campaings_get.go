package endpoints

import (
	"net/http"
)

func (h *HandlerCampaign) CampaignGet(w http.ResponseWriter, r *http.Request) (EndpointStruct, error) {

	campaigns, err := h.CampaignService.Get()

	return EndpointStruct{
		Obj: campaigns, Status: 200,
	}, err

}
