package controllers

import (
	models "URLShortener/Models"
	"URLShortener/dao"
	"encoding/json"
	"net/http"
)

// Estructura para la solicitud de acortar URL
type shortenURLRequest struct {
	OriginalURL string `json:"original_url"`
}

// Estructura para la respuesta de URL acortada
type shortenURLResponse struct {
	ShortenedURL string `json:"shortened_url"`
}

func ShortenURLHandler(w http.ResponseWriter, r *http.Request) {
	// Decodificar la solicitud
	var request shortenURLRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Crear y validar la URL
	url := models.URL{OriginalURL: request.OriginalURL}
	if err := url.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Generar la URL corta pasando la URL original como parámetro
	url.ShortenedURL = url.GenerateShortURL(url.OriginalURL) // Ajustado aquí

	dao, err := dao.NewURLDao()
	if err != nil { //!Err Handling

		return
	}
	err = dao.Save(&url)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Devolver la URL corta
	response := shortenURLResponse{ShortenedURL: url.ShortenedURL}
	json.NewEncoder(w).Encode(response)
}
