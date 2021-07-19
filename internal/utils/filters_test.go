package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func getNums() (slice []int) {
	slice = []int{13, 1, 2, 16, 2, 12, 3, 34, 5, 22, 7, 8, 49}
	return
}

func getStrings() (slice []string) {
	slice = []string{
		"thirteen", "one", "two", "twelve", "six",
		"twenty two", "twenty four", "three",
		"thirty four", "five", "seven", "eight", "forty nine"}
	return
}

func TestFilterEven(t *testing.T) {
	filtered := FilterInt(getNums(), func(i int) bool {
		return i%2 == 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 7)
}

func TestFilterTwos(t *testing.T) {
	filtered := FilterString(getStrings(), func(s string) bool {
		return strings.HasPrefix(s, "tw")
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 4)
}

func TestThreeLetters(t *testing.T) {
	filtered := FilterString(getStrings(), func(s string) bool {
		return len(s) == 3
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 3)
}

func TestFilterOdd(t *testing.T) {
	filtered := FilterInt(getNums(), func(i int) bool {
		return i%2 != 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 6)
}

func TestFilterEmptyResult(t *testing.T) {
	filtered := FilterInt(getNums(), func(i int) bool {
		return i%15 == 0
	})

	assert.NotEqual(t, filtered, nil)
	assert.Equal(t, len(filtered), 0)
}

func TestFilterEmpty(t *testing.T) {
	filteredInts := FilterInt([]int{}, func(i int) bool {
		return i != 0
	})

	filteredStrings := FilterString([]string{}, func(i string) bool {
		return i != ""
	})

	assert.NotEqual(t, filteredInts, nil)
	assert.Equal(t, len(filteredInts), 0)
	assert.NotEqual(t, filteredStrings, nil)
	assert.Equal(t, len(filteredStrings), 0)
}
