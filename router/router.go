package router

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/thomblweed/thomblweed-htmx/views"
)

func Home() chi.Router {
	router := chi.NewRouter()
	router.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))
	router.Get("/", templ.Handler(views.Home()).ServeHTTP)

	return router
}
