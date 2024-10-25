package main

import "testing"

func TestSimulate32(t *testing.T) {
	result := simulate(9, 32)
	if result != 32 {
		t.Errorf("Expected 32 but got %d", result)
	}
}

func TestSimulate10Users(t *testing.T) {
	result := simulate(10, 1618)
	if result != 8317 {
		t.Errorf("Expected 8317 but got %d", result)
	}
}

func TestSimulate13Users(t *testing.T) {
	result := simulate(13, 7999)
	if result != 146373 {
		t.Errorf("Expected 146373 but got %d", result)
	}
}

func TestSimulate30Users(t *testing.T) {
	result := simulate(30, 5807)
	if result != 37305 {
		t.Errorf("Expected 37305 but got %d", result)
	}
}

func TestPart1(t *testing.T) {
	result := simulate(425, 70848)
	if result != 413188 {
		t.Errorf("Expected 413188 but got %d", result)
	}
}
