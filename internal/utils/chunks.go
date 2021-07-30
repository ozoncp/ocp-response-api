package utils

import (
	"fmt"
	"github.com/ozoncp/ocp-response-api/internal/models"
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
		return fmt.Errorf("error while validating input params: chunkSize must be more zero, chunkSize: [%d]", chunkSize)
	}
	if length < chunkSize {
		return fmt.Errorf("error while validating input params: chunkSize must be less than slice size, chunkSize: [%d]", chunkSize)
	}
	return nil
}

// ChunkString makes chunks from slice of strings
func ChunkString(slice []string, chunkSize int) ([][]string, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunksCount := chunkBufferSize(l, chunkSize)
	res := make([][]string, chunksCount)

	start := 0
	i := 0
	for ; i < chunksCount-1; i++ {
		res[i] = slice[start : start+chunkSize]
		start += chunkSize
	}
	res[i] = slice[start:]

	return res, nil
}

// ChunkInt makes chunks from slice of ints
func ChunkInt(slice []int, chunkSize int) ([][]int, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunksCount := chunkBufferSize(l, chunkSize)
	res := make([][]int, chunksCount)

	start := 0
	i := 0
	for ; i < chunksCount-1; i++ {
		res[i] = slice[start : start+chunkSize]
		start += chunkSize
	}
	res[i] = slice[start:]

	return res, nil
}

// ChunkResponse makes chunks from slice of Responses
func ChunkResponse(slice []models.Response, chunkSize int) ([][]models.Response, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunksCount := chunkBufferSize(l, chunkSize)
	res := make([][]models.Response, chunksCount)

	start := 0
	i := 0
	for ; i < chunksCount-1; i++ {
		res[i] = slice[start : start+chunkSize]
		start += chunkSize
	}
	res[i] = slice[start:]

	return res, nil
}
