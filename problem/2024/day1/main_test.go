package main

import "testing"

func TestPart1(t *testing.T) {
	input := []string{"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	total := part1(input)
	if total != 11 {
		t.Errorf("Expected 11, got %d", total)
	}
}

func TestPart2(t *testing.T) {
	input := []string{"3   4",
		"4   3",
		"2   5",
		"1   3",
		"3   9",
		"3   3",
	}
	total := part2(input)
	if total != 31 {
		t.Errorf("Expected 31, got %d", total)
	}
}
