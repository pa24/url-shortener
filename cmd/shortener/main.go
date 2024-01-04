package main

import (
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
	"strings"
)

var urlStore = make(map[string]string)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc(`/`, shortenURLHandler)
	mux.HandleFunc(`/{id}`, redirectHandler)

	err := http.ListenAndServe(`:8080`, mux)
	if err != nil {
		panic(err)
	}
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Error reading request body", http.StatusInternalServerError)
		return
	}

	originalURL := string(body)
	shortURL := generateShortURL()
	urlStore[shortURL] = originalURL
	response := fmt.Sprintf("http://localhost:8080/%s\n", shortURL)

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "text/plain")
	//w.Header().Set("Content-Length", fmt.Sprint(len(response)))

	w.Write([]byte(response))
}

func generateShortURL() string {
	id := uuid.New().String()
	// Remove dashes to make it more URL-friendly
	return strings.ReplaceAll(id, "-", "")
}

func redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "only GET allow", http.StatusMethodNotAllowed)
	}
	id := strings.TrimPrefix(r.URL.Path, "/")

	originalURL, exists := urlStore[id]
	if !exists {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)

}
