package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/thomblweed/thomblweed-htmx/router/routes"
)

func Get() chi.Router {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Handle("/static/*", http.StripPrefix("/static", http.FileServer(http.Dir("./public"))))

	router.Mount("/", routes.Root())
	router.Mount("/blogs", routes.Blogs())

	return router
}
