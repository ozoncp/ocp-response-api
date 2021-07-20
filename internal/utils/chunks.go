package utils

import (
	"errors"
	"github.com/ozoncp/ocp-response-api/domain"
)

func chunkBufferSize(length int, chunkSize int) int {
	bufferSize := length / chunkSize
	if length%chunkSize != 0 {
		bufferSize = bufferSize + 1
	}

	return bufferSize
}

func validate(length int, chunkSize int) error {
	if chunkSize <= 0 {
		return errors.New(ChunkSizeLEqualZeroError)
	}
	if length < chunkSize {
		return errors.New(ChunkSizeGreaterSliceSizeError)
	}
	return nil
}

// ChunkString Makes chunks from slice of strings
func ChunkString(slice []string, chunkSize int) ([][]string, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]string, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}

// ChunkInt Makes chunks from slice of ints
func ChunkInt(slice []int, chunkSize int) ([][]int, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]int, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}

// ChunkResponse Makes chunks from slice of Responses
func ChunkResponse(slice []domain.Response, chunkSize int) ([][]domain.Response, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]domain.Response, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}
