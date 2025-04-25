package routes

import (
	"github.com/go-chi/chi/v5"
	"github.com/jimitchavdadev/url-shortener/internal/handlers"
)

func NewRouter(urlHandler *handlers.URLHandler) *chi.Mux {
	r := chi.NewRouter()

	r.Post("/shorten", urlHandler.ShortenURL)
	r.Get("/{shortCode}", urlHandler.RedirectURL)
	r.Get("/analytics/{shortCode}", urlHandler.GetAnalytics)

	return r
}
