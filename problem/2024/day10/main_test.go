package main

import (
	"strings"
	"testing"
)

func TestPart1Small(t *testing.T) {
	inputTopo := []string{"...0...", "...1...", "...2...", "6543456", "7.....7", "8.....8", "9.....9"}
	expected := 2
	result := part1(inputTopo)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1FourEndsOneHead(t *testing.T) {
	input := "..90..9\n...1.98\n...2..7\n6543456\n765.987\n876....\n987...."
	inputTopo := strings.Split(input, "\n")
	expected := 4
	result := part1(inputTopo)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1TwoHeads(t *testing.T) {
	input := "10..9..\n2...8..\n3...7..\n4567654\n...8..3\n...9..2\n.....01"
	inputTopo := strings.Split(input, "\n")
	expected := 3
	result := part1(inputTopo)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}

func TestPart1Full(t *testing.T) {
	input := "89010123\n78121874\n87430965\n96549874\n45678903\n32019012\n01329801\n10456732"
	inputTopo := strings.Split(input, "\n")
	expected := 36
	result := part1(inputTopo)
	if result != expected {
		t.Fatalf("expected %d, got %d", expected, result)
	}
}
