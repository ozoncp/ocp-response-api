package saver

import (
	"sync"
	"time"

	"github.com/ozoncp/ocp-response-api/internal/flusher"
	"github.com/ozoncp/ocp-response-api/internal/models"
)

// Saver - интерфейс для сохранения Responses
type Saver interface {
	Save(entity models.Response)
	Close()
}

type saver struct {
	capacity   uint
	flusher    flusher.Flusher
	buffer chan models.Response
	wait       *sync.WaitGroup
	state      int8
	timeToLive time.Duration
}

func (s *saver) init() {

	ticker := time.NewTicker(s.timeToLive)
	s.wait.Add(1)

	go func() {
		requests := make([]models.Response, 0, s.capacity)
		defer s.wait.Done()
		for {
			select {
			case req, ok := <-s.buffer:
				if !ok {
					s.flusher.Flush(requests)
					requests = requests[:0]
					return
				} else {
					requests = append(requests, req)
				}
			case <-ticker.C:
				s.flusher.Flush(requests)
				requests = requests[:0]
			}
		}

	}()
}

// Save сохраняет в буфер Response
func (s *saver) Save(response models.Response) {
	s.buffer <- response
}

// Close закрывает буфер сохранения Responses
func (s *saver) Close() {
	close(s.buffer)
	s.wait.Wait()
}

// NewSaver возвращает Saver с поддержкой переодического сохранения
func NewSaver(capacity uint, flusher flusher.Flusher, timeToLive time.Duration) Saver {
	saverEntity := &saver{
		capacity:   capacity,
		flusher:    flusher,
		buffer: make(chan models.Response, capacity),
		wait:       &sync.WaitGroup{},
		timeToLive: timeToLive,
	}

	saverEntity.init()

	return saverEntity
}
