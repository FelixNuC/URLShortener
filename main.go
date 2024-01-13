package main

import (
	"URLShortener/router"
	"net/http"
)

func main() {

	router.SetupRoutes()

	http.ListenAndServe(":8080", nil)
}
