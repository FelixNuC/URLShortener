package models

import (
	"URLShortener/interfaces"
	"time"
)

var _ interfaces.URLInterface = &URL{} // Verifica que *URL implementa URLInterface

type URL struct {
	OriginalURL  string    `json:"original_url"`
	ShortenedURL string    `json:"shortened_url"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

// Aquí implementamos los métodos de la interfaz URLInterface.
func (u *URL) GenerateShortURL() string {
	// Lógica para generar la URL acortada
	return "short_url" // Debes devolver un string como lo indica la interfaz
}

func (u *URL) Validate() error {
	// Lógica para validar la URL original
	// Asegúrate de devolver un error, nil si no hay error.
	return nil
}

func (u *URL) Save() error {
	// Lógica para guardar la URL en la base de datos
	// Devuelve nil o un error.
	return nil
}

func (u *URL) Delete() error {
	// Lógica para eliminar la URL
	// Devuelve nil o un error.
	return nil
}

func (u *URL) Get() (string, error) {
	// Lógica para obtener la URL acortada
	// Devuelve un string y nil o un error.
	return "original_url", nil
}
