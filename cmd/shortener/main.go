package main

import (
	"github.com/pa24/url-shortener/internal/server"
	"github.com/pa24/url-shortener/internal/storage"
)

func main() {
	repo := storage.NewInMemoryRepository()

	err := server.RunServer(":8080", repo)
	if err != nil {
		panic(err)
	}
}
