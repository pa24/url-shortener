package server

import (
	"github.com/pa24/url-shortener/internal/handlers"
	"github.com/pa24/url-shortener/internal/storage"
	"net/http"
)

func RunServer(addr string, repo storage.Repository) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		handlers.ShortenURLHandler(w, r, repo)
	})

	return http.ListenAndServe(addr, mux)
}
