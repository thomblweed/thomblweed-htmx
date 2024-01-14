package router

import (
	"net/http"

	"github.com/go-chi/chi"
)

func Get() chi.Router {
	router := chi.NewRouter()
	router.Get("/", func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("welcome"))
	})

	return router
}