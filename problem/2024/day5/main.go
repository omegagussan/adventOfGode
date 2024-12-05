package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day5/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	rules, updates := parseInput(input)
	r := parseRules(rules)

	total := 0
	for _, v := range updates {
		values := getValues(v)
		if isCorrectlySorted(values, r) {
			total += getMiddleValue(values)
		}
	}
	return total
}

func part2(input string) int {
	rules, updates := parseInput(input)
	r := parseRules(rules)

	total := 0
	for _, v := range updates {
		values := getValues(v)
		if !isCorrectlySorted(values, r) {
			values = sortValues(values, r)
			total += getMiddleValue(values)
		}
	}
	return total
}

func parseInput(input string) (rules, updates []string) {
	parts := strings.Split(input, "\n\n")
	return strings.Split(parts[0], "\n"), strings.Split(parts[1], "\n")
}

func parseRules(rules []string) map[int][]int {
	r := make(map[int][]int)
	for _, rule := range rules {
		s := strings.Split(rule, "|")
		key, _ := strconv.Atoi(s[0])
		val, _ := strconv.Atoi(s[1])
		r[key] = append(r[key], val)
	}
	return r
}

func getValues(v string) []int {
	parts := strings.Split(v, ",")
	values := make([]int, len(parts))
	for i, p := range parts {
		values[i], _ = strconv.Atoi(p)
	}
	return values
}

func isCorrectlySorted(values []int, r map[int][]int) bool {
	for i, v := range values {
		for _, rule := range r[v] {
			if contains(values[:i], rule) {
				return false
			}
		}
	}
	return true
}

func sortValues(values []int, r map[int][]int) []int {
	for !isCorrectlySorted(values, r) {
		for i := 0; i < len(values); i++ {
			for j := i + 1; j < len(values); j++ {
				if contains(r[values[j]], values[i]) {
					values[i], values[j] = values[j], values[i]
				}
			}
		}
	}
	return values
}

func getMiddleValue(values []int) int {
	if len(values) == 0 {
		return 0
	}
	return values[len(values)/2]
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}
