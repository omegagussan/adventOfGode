package main

import (
	"testing"
)

func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		index    int
		expected []int
	}{
		{"RemoveFirstElement", []int{1, 2, 3, 4}, 0, []int{2, 3, 4}},
		{"RemoveMiddleElement", []int{1, 2, 3, 4}, 2, []int{1, 2, 4}},
		{"RemoveLastElement", []int{1, 2, 3, 4}, 3, []int{1, 2, 3}},
		{"RemoveFromSingleElementSlice", []int{1}, 0, []int{}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := remove(tt.input, tt.index)
			if !equal(result, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsDiffLessThan(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		diff     int
		expected bool
	}{
		{"DiffLessThan3", []int{1, 2, 3}, 3, true},
		{"DiffEqualTo3", []int{1, 4, 7}, 3, true},
		{"DiffGreaterThan3", []int{1, 5, 9}, 3, false},
		{"SingleElement", []int{1}, 3, true},
		{"EmptySlice", []int{}, 3, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isDiffLessThan(tt.input, tt.diff)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsStrictlyIncreasing(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"StrictlyIncreasing", []int{1, 2, 3, 4}, true},
		{"NotStrictlyIncreasing", []int{1, 2, 2, 4}, false},
		{"SingleElement", []int{1}, true},
		{"EmptySlice", []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isStrictlyIncreasing(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestIsStrictlyDecreasing(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected bool
	}{
		{"StrictlyDecreasing", []int{4, 3, 2, 1}, true},
		{"NotStrictlyDecreasing", []int{4, 3, 3, 1}, false},
		{"SingleElement", []int{1}, true},
		{"EmptySlice", []int{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := isStrictlyDecreasing(tt.input)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
