package main

import (
	"net/http"

	"github.com/thomblweed/thomblweed-htmx/routes"
)

func main() {
	//	router.Use(middleware.Logger)

	http.ListenAndServe(":3000", routes.Root())
}
