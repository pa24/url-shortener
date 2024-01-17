package handlers

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// MockRepository представляет фейковую реализацию интерфейса Repository для тестирования.
type MockRepository struct {
	saveCalled  bool
	getCalled   bool
	originalURL string
	shortURL    string
	exists      bool
}

func (m *MockRepository) Save(originalURL string) string {
	m.saveCalled = true
	m.originalURL = originalURL
	return m.shortURL
}

func (m *MockRepository) Get(shortURL string) (string, bool) {
	m.getCalled = true
	return m.originalURL, m.exists
}

func TestShortenURLHandler_Post(t *testing.T) {
	// Создаем фейковый объект репозитория
	mockRepo := &MockRepository{}

	// Создаем фейковый объект http.ResponseWriter
	w := httptest.NewRecorder()

	// Создаем фейковый объект http.Request для метода POST
	r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader("https://practicum.yandex.ru/"))

	// Вызываем обработчик
	ShortenURLHandler(w, r, mockRepo)

	// Проверяем, что метод Save был вызван
	assert.True(t, mockRepo.saveCalled, "Expected Save to be called, but it wasn't")

	// Проверяем, что код ответа - StatusCreated
	assert.Equal(t, http.StatusCreated, w.Code, "Expected status code %d, but got %d", http.StatusCreated, w.Code)

	// Проверяем, что заголовок Content-Type установлен на "text/plain"
	assert.Equal(t, "text/plain", w.Header().Get("Content-Type"),
		"Expected Content-Type to be 'text/plain', but got '%s'", w.Header().Get("Content-Type"))

	// Проверяем, что тело ответа содержит корректный URL
	expectedURL := fmt.Sprintf("http://localhost:8080/%s", mockRepo.shortURL)
	if body := w.Body.String(); body != expectedURL {
		t.Errorf("Expected response body to be '%s', but got '%s'", expectedURL, body)
	}

	assert.Equal(t, expectedURL, w.Body.String(),
		"Expected response body to be '%s', but got '%s'", expectedURL, w.Body.String())
}

func TestShortenURLHandler_Get(t *testing.T) {
	// Создаем фейковый объект репозитория
	mockRepo := &MockRepository{
		originalURL: "http://example.com",
		exists:      true,
	}

	// Создаем фейковый объект http.ResponseWriter
	w := httptest.NewRecorder()

	// Создаем фейковый объект http.Request для метода GET
	r := httptest.NewRequest(http.MethodGet, "/shortURL", nil)

	// Вызываем обработчик
	ShortenURLHandler(w, r, mockRepo)

	// Используем библиотеку testify для утверждений

	// Проверяем, что метод Get был вызван
	assert.True(t, mockRepo.getCalled, "Expected Get to be called, but it wasn't")

	// Проверяем, что код ответа - StatusTemporaryRedirect
	assert.Equal(t, http.StatusTemporaryRedirect, w.Code, "Expected status code %d, but got %d", http.StatusTemporaryRedirect, w.Code)

	// Проверяем, что произошло перенаправление на корректный URL
	assert.Equal(t, mockRepo.originalURL, w.Header().Get("Location"), "Expected redirect to '%s', but got '%s'", mockRepo.originalURL, w.Header().Get("Location"))
}
