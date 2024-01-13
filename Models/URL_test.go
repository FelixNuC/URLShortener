package models

import (
	"testing"
	"time"
)

func TestURL_GenerateShortURL(t *testing.T) {
	tests := []struct {
		name        string
		u           *URL
		originalURL string
		want        string
	}{
		{
			name: "Test con URL válida",
			u: &URL{
				CreatedAt: time.Date(2023, 1, 1, 1, 1, 1, 1, time.FixedZone("", 0)),
				ExpiresAt: time.Date(2025, 1, 1, 1, 1, 1, 1, time.FixedZone("", 0)),
			},
			originalURL: "https://itch.io/",
			want:        "ok",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.GenerateShortURL(tt.originalURL); got != tt.want {
				t.Errorf("URL.GenerateShortURL() = %v, want %v", got, tt.want)
			}
		})
	}
}
