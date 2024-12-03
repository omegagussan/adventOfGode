package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day3/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
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
