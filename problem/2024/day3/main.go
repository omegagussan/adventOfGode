package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day3/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input string) int {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	count := 0
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		count += x * y
	}
	return count
}

func part2(input string) int {
	matchMap := getMatchMap(input)

	count := 0
	execute := true
	for i := range input {
		if strings.HasPrefix(input[i:], "do()") {
			execute = true
		} else if strings.HasPrefix(input[i:], "don't()") {
			execute = false
		}

		if execute {
			if values, found := matchMap[i]; found {
				x, _ := strconv.Atoi(values[0])
				y, _ := strconv.Atoi(values[1])
				count += x * y
			}
		}
	}
	return count
}

func getMatchMap(input string) map[int][]string {
	r := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	matches := r.FindAllStringSubmatch(input, -1)
	matchMap := make(map[int][]string)
	for _, match := range matches {
		index := strings.Index(input, match[0])
		matchMap[index] = match[1:]
	}
	return matchMap
}
