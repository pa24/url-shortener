package storage

import (
	"github.com/google/uuid"
	"strings"
)

// Repository интерфейс определяет методы для взаимодействия с хранилищем данных.

type Repository interface {
	Save(originalURL string) string
	Get(shortURL string) (string, bool)
}

// InMemoryRepository представляет хранилище данных в памяти.
type InMemoryRepository struct {
	urlStore map[string]string
}

func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{urlStore: make(map[string]string)}
}

func (r *InMemoryRepository) Save(originalURL string) string {
	shortURL := generateShortURL()
	r.urlStore[shortURL] = originalURL
	return shortURL
}

func (r *InMemoryRepository) Get(shortURL string) (string, bool) {
	originalURL, exists := r.urlStore[shortURL]
	return originalURL, exists
}

func generateShortURL() string {
	id := uuid.New().String()
	strings.ReplaceAll(id, "-", "")
	return id[:6]
}
