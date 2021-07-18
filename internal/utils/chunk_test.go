package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getOneToTenInts() (slice []int) {
	slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return
}

func getOneToTenStrings() (slice []string) {
	slice = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return
}

func TestChunk4From10(t *testing.T) {
	chunksInts, ferr := ChunkInt(getOneToTenInts(), 4)
	chunksStrings, serr := ChunkString(getOneToTenStrings(), 4)

	assert.Equal(t, ferr, nil)
	assert.NotEqual(t, chunksInts, nil)
	assert.Equal(t, len(chunksInts), 3)
	assert.Equal(t, serr, nil)
	assert.NotEqual(t, chunksStrings, nil)
	assert.Equal(t, len(chunksStrings), 3)
}

func TestChunk0From10(t *testing.T) {
	_, ferr := ChunkInt(getOneToTenInts(), 0)
	_, serr := ChunkString(getOneToTenStrings(), 0)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), ChunkSizeLEqualZero)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), ChunkSizeLEqualZero)
}

func TestChunk11From10(t *testing.T) {
	_, ferr := ChunkInt(getOneToTenInts(), 11)
	_, serr := ChunkString(getOneToTenStrings(), 11)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), ChunkSizeLessSliceSize)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), ChunkSizeLessSliceSize)
}

func TestEmptySlice(t *testing.T) {
	_, ferr := ChunkInt([]int{}, 3)
	_, serr := ChunkString([]string{}, 3)

	assert.NotEqual(t, ferr, nil)
	assert.Equal(t, ferr.Error(), ChunkSizeLessSliceSize)
	assert.NotEqual(t, serr, nil)
	assert.Equal(t, serr.Error(), ChunkSizeLessSliceSize)
}
