package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/ozoncp/ocp-response-api/internal/models"
)

func get4Responses() []models.Response {
	return []models.Response{
		{Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
		{Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
		{Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
		{Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
	}
}

func TestGet4Responses(t *testing.T) {
	tests := []struct {
		slice    []models.Response
		expected map[uint64]models.Response
	}{
		{
			get4Responses(),
			map[uint64]models.Response{
				0: {Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
				1: {Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
				2: {Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
				3: {Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
			},
		},
		{
			append(get4Responses(), models.Response{Id: 1, UserId: 1, RequestId: 1, Text: "test1"}),
			nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ToMap(tt.slice)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
