// URLShortener/Router/router.go

package router

import (
	"URLShortener/controllers"
	"net/http"
)

func SetupRoutes() {
	http.HandleFunc("/shorten", controllers.ShortenURLHandler)
	http.HandleFunc("/", controllers.RedirectURLHandler)
}
