package main

import (
	"testing"
)

func TestPart1Simple(t *testing.T) {
	input := "AAAA\nBBCD\nBBCC\nEEEC"
	expected := 140
	result := part1(input)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1Small(t *testing.T) {
	input := "OOOOO\nOXOXO\nOOOOO\nOXOXO\nOOOOO"
	expected := 772
	result := part1(input)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1Full(t *testing.T) {
	input := "RRRRIICCFF\nRRRRIICCCF\nVVRRRCCFFF\nVVRCCCJFFF\nVVVVCJJCFE\nVVIVCCJJEE\nVVIIICJJEE\nMIIIIIJJEE\nMIIISIJEEE\nMMMISSJEEE"
	expected := 1930
	result := part1(input)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1Other(t *testing.T) {
	input := "AAAAAA\nAAABBA\nAAABBA\nABBAAA\nABBAAA\nAAAAAA"
	expected := (6*6-8)*(4*6+16) + 2*(4*8)
	result := part1(input)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
