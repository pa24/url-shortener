package handlers

import (
	"fmt"
	"github.com/pa24/url-shortener/internal/storage"
	"io"
	"net/http"
	"strings"
)

func ShortenURLHandler(w http.ResponseWriter, r *http.Request, repo storage.Repository) {
	if r.Method == http.MethodPost {
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
	} else if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/")

		originalURL, exists := repo.Get(id)
		if !exists {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		http.Redirect(w, r, originalURL, http.StatusTemporaryRedirect)
	} else {
		http.Error(w, "only GET allowed", http.StatusMethodNotAllowed)
		return
	}
}
