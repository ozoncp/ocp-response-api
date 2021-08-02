package flusher

import (
	"github.com/ozoncp/ocp-response-api/internal/models"
	"github.com/ozoncp/ocp-response-api/internal/repo"
	"github.com/ozoncp/ocp-response-api/internal/utils"
)

// Flusher - интерфейс для сброса респонсов в хранилище
type Flusher interface {
	Flush(entities []models.Response) []models.Response
}

// NewFlusher возвращает Flusher с поддержкой батчевого сохранения
func NewFlusher(chunkSize int, repo repo.Repo) Flusher {
	return &flusher{
		chunkSize: chunkSize,
		repo:      repo,
	}
}

type flusher struct {
	chunkSize int
	repo      repo.Repo
}

// Flush сбрасывает респонсы в хранилище
func (f flusher) Flush(entities []models.Response) []models.Response {
	var res []models.Response
	chunks, err := utils.ChunkResponse(entities, 2)

	if err == nil {
		for _, chunk := range chunks {
			if e := f.repo.AddResponses(chunk); e == nil {
				res = append(res, chunk...)
			}
		}
	}

	return res
}
