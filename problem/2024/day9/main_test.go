package main

import (
	"testing"
)

func TestPart1ParseWithValidInput(t *testing.T) {
	input := "12345"
	expected := []Entry{
		{ID: 0, BlockFiles: 1, FreeSpace: 2},
		{ID: 1, BlockFiles: 3, FreeSpace: 4},
		{ID: 2, BlockFiles: 5, FreeSpace: 0},
	}
	result := parseInput(input)
	if len(result) != len(expected) {
		t.Fatalf("expected %d parts, got %d", len(expected), len(result))
	}
	for i, part := range result {
		if part != expected[i] {
			t.Errorf("expected part %v, got %v", expected[i], part)
		}
	}
}

func TestPart1CompressWithValidInput(t *testing.T) {
	input := []Entry{
		{ID: 0, BlockFiles: 1, FreeSpace: 2},
		{ID: 1, BlockFiles: 3, FreeSpace: 4},
		{ID: 2, BlockFiles: 5, FreeSpace: 0},
	}
	expected := []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}
	result := Compress(input)
	if len(result) != len(expected) {
		t.Fatalf("expected %d parts, got %d", len(expected), len(result))
	}
	for i, part := range result {
		if part != expected[i] {
			t.Errorf("expected part %v, got %v on index %v", expected[i], part, i)
		}
	}
}

func TestPart1WithValidInput(t *testing.T) {
	expected := 1928
	result := part1("2333133121414131402")
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2WithValidInput(t *testing.T) {
	expected := 2858
	result := part2("2333133121414131402")
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart2FragmentWithValidInput(t *testing.T) {
	input := parseInput("2333133121414131402")
	expected := []int{0, 0, 9, 9, 2, 1, 1, 1, 7, 7, 7, -1, 4, 4, -1, 3, 3, 3, -1, -1, -1, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, -1, -1, -1, -1, 8, 8, 8, 8, -1, -1}
	result := Fragment(input)
	if len(result) != len(expected) {
		t.Fatalf("expected %d parts, got %d", len(expected), len(result))
	}
	for i, part := range result {
		if part != expected[i] {
			t.Errorf("expected part %v, got %v on index %v", expected[i], part, i)
		}
	}
}
