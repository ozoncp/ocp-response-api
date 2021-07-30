package flusher

import (
	"github.com/ozoncp/ocp-response-api/internal/models"
	"github.com/ozoncp/ocp-response-api/internal/repo"
)

// Flusher - интерфейс для сброса респонсов в хранилище
type Flusher interface {
	Flush(entities []models.Response) []models.Response
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, responseRepo repo.ResponseRepo) Flusher {
	return &flusher{
		chunkSize:    chunkSize,
		responseRepo: responseRepo,
	}
}

type flusher struct {
	chunkSize    int
	responseRepo repo.ResponseRepo
}

// Flush сбрасывает респонсы в хранилище
func (f flusher) Flush(entities []models.Response) []models.Response {
	if err := f.responseRepo.AddResponses(entities); err != nil {
		panic(err)
	}

	return entities
}
