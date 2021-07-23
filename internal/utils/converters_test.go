package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGet10Responses(t *testing.T) {
	tests := []struct {
		slice    []Response
		expected map[uint64]Response
	}{
		{
			Get10Responses(),
			map[uint64]Response{
				0: {Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
				1: {Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
				2: {Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
				3: {Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
				4: {Id: 4, UserId: 4, RequestId: 4, Text: "text4"},
				5: {Id: 5, UserId: 5, RequestId: 5, Text: "text5"},
				6: {Id: 6, UserId: 6, RequestId: 6, Text: "text6"},
				7: {Id: 7, UserId: 7, RequestId: 7, Text: "text7"},
				8: {Id: 8, UserId: 8, RequestId: 8, Text: "text8"},
				9: {Id: 9, UserId: 9, RequestId: 9, Text: "text9"},
			},
		},
		{
			append(Get10Responses(), Response{Id: 1, UserId: 1, RequestId: 1, Text: "test1"}),
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
