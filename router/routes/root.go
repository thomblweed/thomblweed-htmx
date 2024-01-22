package routes

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/thomblweed/thomblweed-htmx/layouts"
)

func Root() chi.Router {
	router := chi.NewRouter()
	router.Get("/", templ.Handler(layouts.MainLayout()).ServeHTTP)

	return router
}
