package endpoints

import "emailn/internal/domain/campaign"

type HandlerCampaign struct {
	CampaignService campaign.Service
}
