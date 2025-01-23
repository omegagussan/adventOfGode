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

func TestGetLastDigit(t *testing.T) {
	res := getLastDigit(123)
	if res != 3 {
		t.Errorf("Expected 3, got %d", res)
	}
}

func TestPart2(t *testing.T) {
	split := []string{
		"1",
		"2",
		"3",
		"2024",
	}
	res := part2(split)
	if res != 23 {
		t.Errorf("Expected 23, got %d", res)
	}
}

func TestPricesAndDiffs(t *testing.T) {
	resP, resD := pricesAndDiffs("123", 10)

	expP := []int{
		3,
		0,
		6,
		5,
		4,
		4,
		6,
		4,
		4,
		2,
	}

	expD := []int{
		0,
		-3,
		6,
		-1,
		-1,
		0,
		2,
		-2,
		0,
		-2,
	}

	if !(equalArr(resP, expP)) {
		t.Errorf("Expected %v, got %v", expP, resP)
	}
	if !(equalArr(resD, expD)) {
		t.Errorf("Expected %v, got %v", expD, resD)
	}

}

func TestBestSequence(t *testing.T) {
	prices := []int{
		3,
		0,
		6,
		5,
		4,
		4,
		6,
		4,
		4,
		2,
	}
	changes := []int{
		-3,
		6,
		-1,
		-1,
		0,
		2,
		-2,
		0,
		-2,
	}
	m := bestSequence(prices, changes)
	best := 0
	bestSeq := sequence{}
	for k, v := range m {
		if v > best {
			best = v
			bestSeq = k
		}
	}
	if !equal(bestSeq, sequence{-1, 0, 2, -2}) {
		t.Errorf("Expected sequence{-1, 0, 2, -2}, got %v", bestSeq)
	}
	if best != 6 {
		t.Errorf("Expected 6, got %d", best)
	}
}

func equal(a, b sequence) bool {
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func equalArr(a, b []int) bool {
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
