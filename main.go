package main

import (
	"URLShortener/router"
	"net/http"
	// Importa otros paquetes necesarios
)

func main() {
	// Configurar las rutas
	router.SetupRoutes()

	// Iniciar el servidor web
	http.ListenAndServe(":8080", nil)
}
