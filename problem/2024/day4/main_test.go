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

func TestIsCross(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		i, j     int
		expected bool
	}{
		{
			name:     "valid cross pattern SAM",
			input:    []string{"S.M", ".A.", "S.M"},
			i:        0,
			j:        0,
			expected: true,
		},
		{
			name:     "invalid cross pattern",
			input:    []string{"M..", ".B.", "..S"},
			i:        0,
			j:        0,
			expected: false,
		},
		{
			name:     "upside-down cross pattern",
			input:    []string{"M.M", ".A.", "S.S"},
			i:        0,
			j:        0,
			expected: true,
		},
		{
			name:     "Side cross pattern",
			input:    []string{"M.S", ".A.", "M.S"},
			i:        0,
			j:        0,
			expected: true,
		},
		{
			name:     "out of bounds",
			input:    []string{"...", "...", "..."},
			i:        3,
			j:        3,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsCross(tt.input, tt.i, tt.j)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}
