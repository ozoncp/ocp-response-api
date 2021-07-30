package utils

import (
	"github.com/ozoncp/ocp-response-api/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func getOneToTenInts() (slice []int) {
	slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	return
}

func getOneToTenStrings() (slice []string) {
	slice = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	return
}

func get10Responses() []models.Response {
	return []models.Response{
		{Id: 0, UserId: 0, RequestId: 0, Text: "text0"},
		{Id: 1, UserId: 1, RequestId: 1, Text: "text1"},
		{Id: 2, UserId: 2, RequestId: 2, Text: "text2"},
		{Id: 3, UserId: 3, RequestId: 3, Text: "text3"},
		{Id: 4, UserId: 4, RequestId: 4, Text: "text4"},
		{Id: 5, UserId: 5, RequestId: 5, Text: "text5"},
		{Id: 6, UserId: 6, RequestId: 6, Text: "text6"},
		{Id: 7, UserId: 7, RequestId: 7, Text: "text7"},
		{Id: 8, UserId: 8, RequestId: 8, Text: "text8"},
		{Id: 9, UserId: 9, RequestId: 9, Text: "text9"},
	}
}

func TestChunkInt(t *testing.T) {
	tests := []struct {
		slice     []int
		chunkSize int
		expected  [][]int
	}{
		{
			getOneToTenInts(),
			4,
			[][]int{{0, 1, 2, 3}, {4, 5, 6, 7}, {8, 9}},
		},
		{
			getOneToTenInts(),
			0,
			nil,
		},
		{
			getOneToTenInts(),
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
			getOneToTenStrings(),
			4,
			[][]string{{"zero", "one", "two", "three"}, {"four", "five", "six", "seven"}, {"eight", "nine"}},
		},
		{
			getOneToTenStrings(),
			0,
			nil,
		},
		{
			getOneToTenStrings(),
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
		slice     []models.Response
		chunkSize int
		expected  [][]models.Response
	}{
		{
			get10Responses(),
			4,
			[][]models.Response{
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
			get10Responses(),
			0,
			nil,
		},
		{
			get10Responses(),
			11,
			nil,
		},
		{
			[]models.Response{},
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
