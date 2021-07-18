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
		return errors.New(ChunkSizeLEqualZero)
	}
	if length < chunkSize {
		return errors.New(ChunkSizeLessSliceSize)
	}
	return nil
}

// ChunkString Makes chunks of input slice of strings
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

// ChunkInt Makes chunks of input slice of ints
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
