package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFilterInt(t *testing.T) {
	tests := []struct {
		slice    []int
		filter   func(i int) bool
		expected []int
	}{
		{
			slice:    GetNums(),
			filter:   func(i int) bool { return i%2 == 0 },
			expected: []int{2, 16, 2, 12, 34, 22, 8},
		},
		{
			slice:    GetNums(),
			filter:   func(i int) bool { return i%2 != 0 },
			expected: []int{13, 1, 3, 5, 7, 49},
		},
		{
			slice:    GetNums(),
			filter:   func(i int) bool { return i%15 == 0 },
			expected: []int(nil),
		},
		{
			slice:    []int{},
			filter:   func(i int) bool { return i != 0 },
			expected: []int(nil),
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := FilterInt(tt.slice, tt.filter)
			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestFilterStrings(t *testing.T) {
	tests := []struct {
		slice    []string
		filter   func(i string) bool
		expected []string
	}{
		{
			slice:    GetStrings(),
			filter:   func(s string) bool { return strings.HasPrefix(s, "tw") },
			expected: []string{"two", "twelve", "twenty two", "twenty four"},
		},
		{
			slice:    GetStrings(),
			filter:   func(s string) bool { return len(s) == 3 },
			expected: []string{"one", "two", "six"},
		},
		{
			slice:    []string{},
			filter:   func(s string) bool { return s != "" },
			expected: []string(nil),
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual := FilterString(tt.slice, tt.filter)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
