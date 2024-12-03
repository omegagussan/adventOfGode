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
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	m := r.FindAllStringSubmatch(input, -1)
	count := 0
	for _, v := range m {
		a, b := v[1], v[2]
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)
		count += x * y
	}
	return count
}

func part2(input string) int {
	r := regexp.MustCompile("mul\\((\\d+),(\\d+)\\)")
	m := r.FindAllStringSubmatch(input, -1)
	matches := make(map[int][]string)
	for _, v := range m {
		//find first occurrence of mul
		i := strings.Index(input, v[0])
		matches[i] = v[1:]
	}

	count := 0
	e := true
	for i, _ := range input {
		if strings.HasPrefix(input[i:], "do()") {
			e = true
		} else if strings.HasPrefix(input[i:], "don't()") {
			e = false
		}

		if e {
			if v, ok := matches[i]; ok {
				x, _ := strconv.Atoi(v[0])
				y, _ := strconv.Atoi(v[1])
				count += x * y
			}
		}
	}
	return count
}

//88811886 (all operations are unique)
