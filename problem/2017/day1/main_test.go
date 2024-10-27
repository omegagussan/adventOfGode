package main

import (
	"testing"
)

func TestGetSumWithConsecutiveDuplicates(t *testing.T) {
	arr := []int{1, 1, 2, 2}
	expected := 3
	result := getSum(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetSumWithWrappaAroundDupplicates(t *testing.T) {
	arr := []int{9, 1, 2, 1, 2, 1, 2, 9}
	expected := 9
	result := getSum(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetSumWithNoDuplicates(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := 0
	result := getSum(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetSumWithAllDuplicates(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	expected := 4
	result := getSum(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestGetSumWithEmptyArray(t *testing.T) {
	var arr []int
	expected := 0
	result := getSum(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumPart2WithConsecutiveDuplicates(t *testing.T) {
	arr := []int{1, 2, 1, 2}
	expected := 6
	result := getSumPart2(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumPart2WithNoDuplicates(t *testing.T) {
	arr := []int{1, 2, 3, 4}
	expected := 0
	result := getSumPart2(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumPart2WithAllDuplicates(t *testing.T) {
	arr := []int{1, 1, 1, 1}
	expected := 4
	result := getSumPart2(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumPart2WithEmptyArray(t *testing.T) {
	var arr []int
	expected := 0
	result := getSumPart2(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestSumPart2WithWrapAroundDuplicates(t *testing.T) {
	arr := []int{1, 2, 3, 1, 2, 3}
	expected := 12
	result := getSumPart2(arr)
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}
