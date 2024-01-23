package server

import (
	"github.com/go-chi/chi/v5"
	"github.com/pa24/url-shortener/internal/handlers"
	"github.com/pa24/url-shortener/internal/storage"
	"net/http"
)

func RunServer(addr string, repo storage.Repository) error {
	r := chi.NewRouter()
	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodGet && r.Method != http.MethodPost {
				http.Error(w, "Bad Request", http.StatusBadRequest)
				return
			}
			next.ServeHTTP(w, r)
		})
	})

	r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.RedirectHandler(w, r, repo)
	})
	r.Post("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ShortenURLHandler(w, r, repo)
	})

	return http.ListenAndServe(addr, r)
}
