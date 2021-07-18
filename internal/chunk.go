package utils

import "errors"

func chunkBufferSize(length int, chunkSize int) int {
	bufferSize := length / chunkSize
	if length%chunkSize != 0 {
		bufferSize = bufferSize + 1
	}

	return bufferSize
}

func validate(length int, chunkSize int) error {
	if chunkSize <= 0 {
		return errors.New("Chunk size must be greater than 0")
	}
	if length < chunkSize {
		return errors.New("Chunk size must be less or equal to slice length")
	}
	return nil
}

func Chunk(slice []interface{}, chunkSize int) ([][]interface{}, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]interface{}, chunkBufferSize(l, chunkSize))
	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}

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

func ChunkByte(slice []byte, chunkSize int) ([][]byte, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]byte, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}

func ChunkFloat32(slice []float32, chunkSize int) ([][]float32, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]float32, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}

func ChunkFloat64(slice []float64, chunkSize int) ([][]float64, error) {
	l := len(slice)
	if err := validate(l, chunkSize); err != nil {
		return nil, err
	}

	chunks := make([][]float64, chunkBufferSize(l, chunkSize))

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks, nil
}
