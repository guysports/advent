package day2

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ProcessKey(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Occurences
	}{
		{
			name:  "Test key with no occurences",
			input: "abcde",
			expected: Occurences{
				two:   0,
				three: 0,
			},
		},
		{
			name:  "Test key with a single two occurence",
			input: "abcbe",
			expected: Occurences{
				two:   1,
				three: 0,
			},
		},
		{
			name:  "Test key with a single three occurence",
			input: "abcbb",
			expected: Occurences{
				two:   0,
				three: 1,
			},
		},
		{
			name:  "Test key with a single two and three occurence",
			input: "abcbbia",
			expected: Occurences{
				two:   1,
				three: 1,
			},
		},
		{
			name:  "Test key with multiple three occurences",
			input: "abcbbiaa",
			expected: Occurences{
				two:   0,
				three: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ProcessKey(tt.input)
			assert.Equal(t, got, &tt.expected)

		})
	}
}

func Test_CreateChecksum(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected int
	}{
		{
			name:     "Process inputs with no occurences",
			input:    []string{"abcde", "fghij"},
			expected: 0,
		},
		{
			name:     "Process inputs with one two and one three occurences",
			input:    []string{"abcbe", "fgggj"},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CreateChecksum(tt.input)
			assert.Equal(t, got, tt.expected)

		})
	}
}
