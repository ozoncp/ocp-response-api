package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseIntInt(t *testing.T) {
	tests := []struct {
		dict     map[int]int
		expected map[int]int
	}{
		{
			dict:     map[int]int{0: 1, 1: 2},
			expected: map[int]int{1: 0, 2: 1},
		},
		{
			dict:     map[int]int{},
			expected: map[int]int{},
		},
		{
			dict:     map[int]int{0: 1, 1: 1},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ReverseIntInt(tt.dict)
			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestReverseIntString(t *testing.T) {
	tests := []struct {
		dict     map[int]string
		expected map[string]int
	}{
		{
			dict:     map[int]string{0: "1", 1: "2"},
			expected: map[string]int{"1": 0, "2": 1},
		},
		{
			dict:     map[int]string{},
			expected: map[string]int{},
		},
		{
			dict:     map[int]string{0: "1", 1: "1"},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ReverseIntString(tt.dict)
			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestReverseStringInt(t *testing.T) {
	tests := []struct {
		dict     map[string]int
		expected map[int]string
	}{
		{
			dict:     map[string]int{"0": 1, "1": 2},
			expected: map[int]string{1: "0", 2: "1"},
		},
		{
			dict:     map[string]int{},
			expected: map[int]string{},
		},
		{
			dict:     map[string]int{"0": 1, "1": 1},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ReverseStringInt(tt.dict)
			assert.Equal(t, actual, tt.expected)
		})
	}
}

func TestReverseStringString(t *testing.T) {
	tests := []struct {
		dict     map[string]string
		expected map[string]string
	}{
		{
			dict:     map[string]string{"0": "1", "1": "2"},
			expected: map[string]string{"1": "0", "2": "1"},
		},
		{
			dict:     map[string]string{},
			expected: map[string]string{},
		},
		{
			dict:     map[string]string{"0": "1", "1": "1"},
			expected: nil,
		},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			actual, _ := ReverseStringString(tt.dict)
			assert.Equal(t, actual, tt.expected)
		})
	}
}
