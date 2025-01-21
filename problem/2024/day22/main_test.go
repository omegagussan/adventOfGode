package main

import (
	"testing"
)

func TestPrune(t *testing.T) {
	res := prune(100000000)
	if res != 16113920 {
		t.Errorf("Expected 16113920, got %d", res)
	}
}

func TestNextOneStep(t *testing.T) {
	res := next(123)
	if res != 15887950 {
		t.Errorf("Expected 15887950, got %d", res)
	}
}

func TestNextTenSteps(t *testing.T) {
	expected := []int{
		15887950,
		16495136,
		527345,
		704524,
		1553684,
		12683156,
		11100544,
		12249484,
		7753432,
		5908254,
	}
	results := make([]int, 10)
	res := 123
	for i := 0; i < 10; i++ {
		res = next(res)
		results[i] = res
	}
	for i, r := range results {
		if r != expected[i] {
			t.Errorf("Expected %d, got %d", expected[i], r)
		}
	}
}

func TestPart1(t *testing.T) {
	split := []string{
		"1",
		"10",
		"100",
		"2024",
	}
	res := part1(split)
	if res != 37327623 {
		t.Errorf("Expected 37327623, got %d", res)
	}
}
