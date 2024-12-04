package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "rotate 2x2 matrix",
			input:    []string{"ab", "cd"},
			expected: []string{"ca", "db"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := RotateClockwise(tt.input)
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("expected %v, got %v", tt.expected, result)
				}
			}
		})
	}
}
