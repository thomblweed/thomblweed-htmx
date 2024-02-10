package routes

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	features "github.com/thomblweed/thomblweed-htmx/features"
)

func Blogs() chi.Router {
	router := chi.NewRouter()
	router.Get("/", templ.Handler(features.Blogs()).ServeHTTP)

	return router
}
