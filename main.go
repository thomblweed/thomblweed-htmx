package main

import (
	"net/http"

	router "github.com/thomblweed/thomblweed-htmx/router"
)

func main() {
	http.ListenAndServe(":3000", router.Home())
}
