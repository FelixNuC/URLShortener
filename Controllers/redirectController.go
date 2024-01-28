package controllers

import (
	"URLShortener/DAO"
	"net/http"
)

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {

	shortURL := r.URL.Path[len("/"):]

	dao, err := DAO.NewURLDao()
	if err != nil {
		http.Error(w, "Error interno", http.StatusInternalServerError)
		return
	}

	urlModel, err := dao.Get(shortURL)
	if err != nil {

		http.Error(w, "URL no encontrada", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, urlModel.OriginalURL, http.StatusFound)
}
