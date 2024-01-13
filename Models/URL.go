package models

import (
	"URLShortener/interfaces"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"net/url"
	"time"
)

var _ interfaces.URLInterface = &URL{}

type URL struct {
	OriginalURL  string    `json:"original_url"`
	ShortenedURL string    `json:"shortened_url"`
	CreatedAt    time.Time `json:"created_at"`
	ExpiresAt    time.Time `json:"expires_at"`
}

func (u *URL) GenerateShortURL(originalURL string) string {
	Ourl := u.OriginalURL
	if originalURL == "" {

		return ""
	}
	hasher := sha256.New()
	hasher.Write([]byte(Ourl))
	sha := hasher.Sum(nil)

	shortURL := base64.URLEncoding.EncodeToString(sha[:])

	if len(shortURL) > 10 {
		shortURL = shortURL[:10]
	}

	return shortURL
}

func (u *URL) Validate() error {

	if u.OriginalURL == "" {
		return errors.New("la URL original no puede estar vacía")
	}

	parsedURL, err := url.ParseRequestURI(u.OriginalURL)
	if err != nil {
		return errors.New("la URL original no es válida")
	}

	if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return errors.New("la URL debe comenzar con http o https")
	}

	return nil
}
