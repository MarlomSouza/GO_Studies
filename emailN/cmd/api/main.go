package main

import (
	"emailn/internal/contract"
	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/{name}", func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, "name")
		w.Write([]byte(param))
	})

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		queryParam := r.URL.Query().Get("id")
		w.Write([]byte(queryParam))

	})

	r.Get("/json", func(w http.ResponseWriter, r *http.Request) {

		obj := map[string]string{"message": "success"}
		render.JSON(w, r, obj)

	})

	r.Get("/campaign", func(w http.ResponseWriter, r *http.Request) {
		println("running api route")
		campaign := contract.NewCampaignDto{}
		render.DecodeJSON(r.Body, campaign)
		render.JSON(w, r, campaign)
	})

	http.ListenAndServe(":3000", r)
}

func myMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		println(r.Method, r.RequestURI)
		next.ServeHTTP(w, r)
		println("Running myMiddleware after")
	})
}
