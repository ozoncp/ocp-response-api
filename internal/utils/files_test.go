package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	file            string
	repeats         int
	expectedRepeats int
	isError         bool
}

func getFileRepeatsData() []testCase {
	return []testCase{
		{"chunks.go", 10000, 10000, false},
		{"chunks_2.go", 1, 0, true},
		{"", 1, 0, true},
		{"chunks.go", 0, 0, false},
	}
}

func TestRepeatableReadFiles(t *testing.T) {
	tests := getFileRepeatsData()

	for _, item := range tests {
		result, err := RepeatableRead(item.file, item.repeats)

		if item.isError {
			assert.NotEqual(t, err, nil)
		} else {
			assert.NotEqual(t, result, nil)
			assert.Equal(t, len(result), item.expectedRepeats)
		}
	}
}
