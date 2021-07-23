package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestChunkInt(t *testing.T) {
	tests := []struct {
		slice     []int
		chunkSize int
		expected  [][]int
	}{
		{
			GetOneToTenInts(),
			4,
			[][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			GetOneToTenInts(),
			0,
			nil,
		},
		{
			GetOneToTenInts(),
			11,
			nil,
		},
		{
			[]int{},
			3,
			nil,
		},
	}

	for _, tcase := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ChunkInt(tcase.slice, tcase.chunkSize)
			assert.ElementsMatch(t, actual, tcase.expected)
		})
	}
}

func TestChunkStrings(t *testing.T) {
	tests := []struct {
		slice     []string
		chunkSize int
		expected  [][]string
	}{
		{
			GetOneToTenStrings(),
			4,
			[][]string{{"zero", "one", "two", "three"}, {"four", "five", "six", "seven"}, {"eight", "nine"}},
		},
		{
			GetOneToTenStrings(),
			0,
			nil,
		},
		{
			GetOneToTenStrings(),
			11,
			nil,
		},
		{
			[]string{},
			3,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ChunkString(tt.slice, tt.chunkSize)
			assert.ElementsMatch(t, actual, tt.expected)
		})
	}
}

func TestChunkResponses(t *testing.T) {
	tests := []struct {
		slice     []Response
		chunkSize int
		expected  [][]Response
	}{
		{
			Get10Responses(),
			4,
			[][]Response{
				{
					{Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
					{Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
					{Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
					{Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
				},
				{
					{Id: 4, UserId: 4, RequestId: 4, Text: "text4"},
					{Id: 5, UserId: 5, RequestId: 5, Text: "text5"},
					{Id: 6, UserId: 6, RequestId: 6, Text: "text6"},
					{Id: 7, UserId: 7, RequestId: 7, Text: "text7"},
				},
				{
					{Id: 8, UserId: 8, RequestId: 8, Text: "text8"},
					{Id: 9, UserId: 9, RequestId: 9, Text: "text9"},
				},
			},
		},
		{
			Get10Responses(),
			0,
			nil,
		},
		{
			Get10Responses(),
			11,
			nil,
		},
		{
			[]Response{},
			3,
			nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ChunkResponse(tt.slice, tt.chunkSize)
			assert.ElementsMatch(t, actual, tt.expected)
		})
	}
}
