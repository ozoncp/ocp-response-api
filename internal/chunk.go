package utils

func Chunk(slice []interface{}, chunkSize int) [][]interface{} {
	l := len(slice)
	chunks := make([][]interface{}, l/chunkSize)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}

func ChunkString(slice []string, chunkSize int) [][]string {
	l := len(slice)
	chunks := make([][]string, l/chunkSize+1)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}

func ChunkInt(slice []int, chunkSize int) [][]int {
	l := len(slice)
	chunks := make([][]int, l/chunkSize+1)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}

func ChunkByte(slice []byte, chunkSize int) [][]byte {
	l := len(slice)
	chunks := make([][]byte, l/chunkSize+1)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}

func ChunkFloat32(slice []float32, chunkSize int) [][]float32 {
	l := len(slice)
	chunks := make([][]float32, l/chunkSize+1)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}

func ChunkFloat64(slice []float64, chunkSize int) [][]float64 {
	l := len(slice)
	chunks := make([][]float64, l/chunkSize+1)

	for i := 0; i < l; i++ {
		chunks[i/chunkSize] = append(chunks[i/chunkSize], slice[i])
	}

	return chunks
}
