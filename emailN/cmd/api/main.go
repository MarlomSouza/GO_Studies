package main

import (
	"emailn/internal/domain/campaign"
	"emailn/internal/endpoints"
	"emailn/internal/infrastructure/database"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	db := database.NewDb()
	repository := database.CampaignRepository{Db: db}
	service := campaign.ServiceImp{Repository: &repository}
	handler := endpoints.HandlerCampaign{
		CampaignService: &service,
	}

	r.Post("/campaigns", endpoints.HandlerError(handler.CampaignPost))
	r.Get("/campaigns", endpoints.HandlerError(handler.CampaignGet))
	r.Get("/campaigns/{id}", endpoints.HandlerError(handler.CampaignGetById))
	r.Patch("/campaigns/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
	r.Delete("/campaigns/{id}", endpoints.HandlerError(handler.CampaignDelete))

	http.ListenAndServe(":3000", r)
}
