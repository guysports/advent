package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Inputs(t *testing.T) {
	tests := []struct {
		name          string
		input         string
		expected      int
		expectedError bool
	}{
		{
			name:          "Valid positive input",
			input:         "+1",
			expected:      1,
			expectedError: false,
		},
		{
			name:          "Valid negative input",
			input:         "-1",
			expected:      -1,
			expectedError: false,
		},
		{
			name:          "Invalid prefix",
			input:         "10",
			expectedError: true,
		},
		{
			name:          "Invalid number",
			input:         "+1a",
			expectedError: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			modifier, err := validateInput(tt.input)
			if tt.expectedError {
				assert.NotNil(t, err)
			} else {
				assert.Nil(t, err)
				assert.Equal(t, modifier, tt.expected)
			}
		})
	}
}

func Test_ProcessInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected Frequency
	}{
		{
			name:  "Test valid input",
			input: "+2,-1,+3,-2,+4,-3",
			expected: Frequency{
				freq:       2,
				historic:   []int{2, 1, 4, 2},
				repeatFreq: 2,
			},
		},
		{
			name:  "Test input with invalid entry",
			input: "+2,-1,+3,-2,2,+4,-3",
			expected: Frequency{
				freq:       2,
				historic:   []int{2, 1, 4, 2},
				repeatFreq: 2,
			},
		},
		{
			name:  "Test input with find historic frequency repeat",
			input: "+1,-1",
			expected: Frequency{
				freq:       1,
				historic:   []int{1, 0, 1},
				repeatFreq: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			frequency := Frequency{
				freq: 0,
			}
			frequency.ProcessFrequencyInput(tt.input)
			assert.Equal(t, frequency, tt.expected)
		})
	}
}
