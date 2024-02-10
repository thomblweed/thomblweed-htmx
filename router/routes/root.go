package routes

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	features "github.com/thomblweed/thomblweed-htmx/features"
	"github.com/thomblweed/thomblweed-htmx/layouts"
)

func Root() chi.Router {
	router := chi.NewRouter()
	router.Get("/", handleNavigationBoost)

	return router
}

func handleNavigationBoost(w http.ResponseWriter, r *http.Request) {
	isBoosted := r.Header.Get("HX-Boosted") == "true"
	if isBoosted {
		templ.Handler(features.AboutMe()).ServeHTTP(w, r)
		return
	}

	templ.Handler(layouts.MainLayout(features.AboutMe)).ServeHTTP(w, r)
}
