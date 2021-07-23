package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRepeatableReadFiles(t *testing.T) {
	tests := GetFileRepeatsData()

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
