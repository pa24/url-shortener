package util

import (
	"github.com/google/uuid"
	"strings"
)

func GenerateShortURL() string {
	id := uuid.New().String()
	strings.ReplaceAll(id, "-", "")
	return id[:6]
}
