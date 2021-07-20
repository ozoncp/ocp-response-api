package tests

import (
	"github.com/ozoncp/ocp-response-api/domain"
	"github.com/ozoncp/ocp-response-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestChunk4From10(t *testing.T) {
	chunksInts, ferr := utils.ChunkInt(GetOneToTenInts(), 4)
	chunksStrings, serr := utils.ChunkString(GetOneToTenStrings(), 4)

	assert.Equal(t, ferr, nil)
	assert.NotEqual(t, chunksInts, nil)
	assert.Equal(t, len(chunksInts), 3)
	assert.Equal(t, serr, nil)
	assert.NotEqual(t, chunksStrings, nil)
	assert.Equal(t, len(chunksStrings), 3)
}

func TestChunk0From10(t *testing.T) {
	_, ferr := utils.ChunkInt(GetOneToTenInts(), 0)
	_, serr := utils.ChunkString(GetOneToTenStrings(), 0)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), utils.ChunkSizeLEqualZeroError)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), utils.ChunkSizeLEqualZeroError)
}

func TestChunk11From10(t *testing.T) {
	_, ferr := utils.ChunkInt(GetOneToTenInts(), 11)
	_, serr := utils.ChunkString(GetOneToTenStrings(), 11)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), utils.ChunkSizeGreaterSliceSizeError)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), utils.ChunkSizeGreaterSliceSizeError)
}

func TestEmptySlice(t *testing.T) {
	_, ferr := utils.ChunkInt([]int{}, 3)
	_, serr := utils.ChunkString([]string{}, 3)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), utils.ChunkSizeGreaterSliceSizeError)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), utils.ChunkSizeGreaterSliceSizeError)
}

func TestChunk3ResponseFrom10(t *testing.T) {
	responses := Get10Responses()

	chunks, err := utils.ChunkResponse(responses, 3)

	assert.Equal(t, err, nil)
	assert.NotEqual(t, chunks, nil)
	assert.Equal(t, len(chunks), 4)
}

func TestChunk0ResponseFrom10(t *testing.T) {
	responses := Get10Responses()
	_, err := utils.ChunkResponse(responses, 0)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), utils.ChunkSizeLEqualZeroError)
}

func TestChunk11ResponseFrom10(t *testing.T) {
	responses := Get10Responses()
	_, err := utils.ChunkResponse(responses, 11)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), utils.ChunkSizeGreaterSliceSizeError)
}

func TestEmptyResponseSlice(t *testing.T) {
	_, err := utils.ChunkResponse([]domain.Response{}, 3)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, err.Error(), utils.ChunkSizeGreaterSliceSizeError)
}
