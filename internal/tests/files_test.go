package tests

import (
	"github.com/ozoncp/ocp-response-api/internal/utils"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRepeatableReadFiles(t *testing.T) {
	tests := GetFileRepeatsData()

	for _, item := range tests {
		result, err := utils.RepeatableRead(item.file, item.repeats)

		if item.isError {
			assert.NotEqual(t, err, nil)
		} else {
			assert.NotEqual(t, result, nil)
			assert.Equal(t, len(result), item.expectedRepeats)
		}
	}
}
