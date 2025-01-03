package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day19/input.txt")
	towels, patterns := parse(string(bytes))
	fmt.Println(part1(towels, patterns))
	fmt.Println(part2(towels, patterns))
}

func parse(input string) ([]string, []string) {
	parts := strings.Split(input, "\n\n")
	return strings.Split(parts[0], ", "), strings.Split(parts[1], "\n")
}

func part1(towels []string, patterns []string) int {
	count := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		if solutions(pattern, towels, cache) != 0 {
			count++
		}
	}
	return count
}

func part2(towels []string, patterns []string) int {
	count := 0
	cache := make(map[string]int)
	for _, pattern := range patterns {
		count += solutions(pattern, towels, cache)
	}
	return count
}

func solutions(pattern string, towels []string, cache map[string]int) int {
	if val, ok := cache[pattern]; ok {
		return val
	}
	if pattern == "" {
		return 1
	}
	acc := 0
	for _, towel := range towels {
		if strings.HasPrefix(pattern, towel) {
			acc += solutions(pattern[len(towel):], towels, cache)
		}
	}
	cache[pattern] = acc
	return acc
}
