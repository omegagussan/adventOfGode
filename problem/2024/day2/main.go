package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day2/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println(part1(split))
}

func part1(split []string) int {
	safe := 0
	for _, s := range split {
		numbers := make([]int, 0)
		vs := strings.Split(s, " ")
		for _, v := range vs {
			numb, _ := strconv.Atoi(v)
			numbers = append(numbers, numb)
		}
		if (isStrictlyIncreasing(numbers) || isStrictlyDecreasing(numbers)) &&
			isDiffLessThan(numbers, 3) {
			safe++
		}

	}
	return safe
}

func isDiffLessThan(numbers []int, diff int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		var abs = math.Abs(float64(numbers[i+1] - numbers[i]))
		if abs > float64(diff) {
			return false
		}
	}
	return true
}

func isStrictlyIncreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if !(numbers[i] < numbers[i+1]) {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if !(numbers[i] > numbers[i+1]) {
			return false
		}
	}
	return true
}
