package tests

import (
	"github.com/ozoncp/ocp-response-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestFilterEven(t *testing.T) {
	filtered := utils.FilterInt(GetNums(), func(i int) bool {
		return i%2 == 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 7)
}

func TestFilterTwos(t *testing.T) {
	filtered := utils.FilterString(GetStrings(), func(s string) bool {
		return strings.HasPrefix(s, "tw")
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 4)
}

func TestThreeLetters(t *testing.T) {
	filtered := utils.FilterString(GetStrings(), func(s string) bool {
		return len(s) == 3
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 3)
}

func TestFilterOdd(t *testing.T) {
	filtered := utils.FilterInt(GetNums(), func(i int) bool {
		return i%2 != 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 6)
}

func TestFilterEmptyResult(t *testing.T) {
	filtered := utils.FilterInt(GetNums(), func(i int) bool {
		return i%15 == 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 0)
}

func TestFilterEmpty(t *testing.T) {
	filteredInts := utils.FilterInt([]int{}, func(i int) bool {
		return i != 0
	})

	filteredStrings := utils.FilterString([]string{}, func(i string) bool {
		return i != ""
	})

	assert.NotEqual(t, filteredInts, nil)
	assert.Equal(t, len(filteredInts), 0)
	assert.NotEqual(t, filteredStrings, nil)
	assert.Equal(t, len(filteredStrings), 0)
}
