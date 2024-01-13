package controllers

import (
	models "URLShortener/Models"
	"URLShortener/dao"
	"encoding/json"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type shortenURLRequest struct {
	OriginalURL string `json:"original_url"`
}
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

	url := models.URL{OriginalURL: request.OriginalURL}
	if err := url.Validate(); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	url.ShortenedURL = url.GenerateShortURL(url.OriginalURL)
	url.CreatedAt = time.Now()
	url.ExpiresAt = time.Now().AddDate(10, 0, 0)
	dao, err := dao.NewURLDao()
	if err != nil { //!Err Handling
		logrus.Error("error al instanciar el dao")
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
