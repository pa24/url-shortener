package handlers

import (
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/pa24/url-shortener/internal/storage"
	"io"
	"net/http"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request, repo storage.Repository) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	originalURL := string(body)
	shortURL := repo.Save(originalURL)
	response := fmt.Sprintf("http://localhost:8080/%s", shortURL)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(response))
}

func RedirectHandler(w http.ResponseWriter, r *http.Request, repo storage.Repository) {
	id := chi.URLParam(r, "id")
	originalURL, exists := repo.Get(id)

	if !exists {
		http.Error(w, "Not Found", http.StatusBadRequest)
		return
	}
	http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)

}
