package controllers

import (
	"URLShortener/DAO"
	"net/http"
)

func RedirectURLHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener el identificador de la URL corta desde la ruta
	shortURL := r.URL.Path[len("/"):]

	// Obtener el DAO y buscar la URL original correspondiente
	dao, err := DAO.NewURLDao()
	if err != nil { //!Err Handling

		return
	}
	url, err := dao.Get(shortURL)
	if err != nil {
		// Si la URL no se encuentra, devuelve un error 404
		http.Error(w, "URL no encontrada", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, url.OriginalURL, http.StatusFound)
}
