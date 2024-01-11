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
	mux.HandleFunc("/", shortenURLHandler)

	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func shortenURLHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusInternalServerError)
			return
		}

		originalURL := string(body)
		shortURL := generateShortURL()
		urlStore[shortURL] = originalURL
		response := fmt.Sprintf("http://localhost:8080/%s", shortURL)

		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte(response))

	} else if r.Method == http.MethodGet {
		id := strings.TrimPrefix(r.URL.Path, "/")

		originalURL, exists := urlStore[id]
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

func generateShortURL() string {
	id := uuid.New().String()
	strings.ReplaceAll(id, "-", "")
	return id[:6]
}
