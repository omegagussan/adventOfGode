package main

import (
	"fmt"
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
	fmt.Println(part2(split))
}

func part1(split []string) int {
	safe := 0
	for _, s := range split {
		numbers := parseNumbers(s)
		if (isStrictlyIncreasing(numbers) || isStrictlyDecreasing(numbers)) && isDiffLessThan(numbers, 3) {
			safe++
		}
	}
	return safe
}

func part2(split []string) int {
	safe := 0
Outer:
	for _, s := range split {
		numbers := parseNumbers(s)
		if (isStrictlyIncreasing(numbers) || isStrictlyDecreasing(numbers)) && isDiffLessThan(numbers, 3) {
			safe++
			continue Outer
		}
		for i := 0; i < len(numbers); i++ {
			tmp := remove(numbers, i)
			if (isStrictlyIncreasing(tmp) || isStrictlyDecreasing(tmp)) && isDiffLessThan(tmp, 3) {
				safe++
				continue Outer
			}
		}
	}
	return safe
}

func parseNumbers(s string) []int {
	vs := strings.Split(s, " ")
	numbers := make([]int, len(vs))
	for i, v := range vs {
		numbers[i], _ = strconv.Atoi(v)
	}
	return numbers
}

func remove(ints []int, i int) []int {
	slice := append([]int(nil), ints...)
	return append(slice[:i], ints[i+1:]...)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i

}

func isDiffLessThan(numbers []int, diff int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if abs(numbers[i+1]-numbers[i]) > diff {
			return false
		}
	}
	return true
}

func isStrictlyIncreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] >= numbers[i+1] {
			return false
		}
	}
	return true
}

func isStrictlyDecreasing(numbers []int) bool {
	for i := 0; i < len(numbers)-1; i++ {
		if numbers[i] <= numbers[i+1] {
			return false
		}
	}
	return true
}
