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
	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	r.Route("/campaigns", func(r chi.Router) {
		r.Use(endpoints.Auth)
		r.Post("/", endpoints.HandlerError(handler.CampaignPost))
		r.Get("/", endpoints.HandlerError(handler.CampaignGet))
		r.Get("/{id}", endpoints.HandlerError(handler.CampaignGetById))
		r.Patch("/cancel/{id}", endpoints.HandlerError(handler.CampaignCancelPatch))
		r.Delete("/{id}", endpoints.HandlerError(handler.CampaignDelete))
	})

	http.ListenAndServe(":3000", r)
}
