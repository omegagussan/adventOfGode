package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Equation struct {
	terms  []int
	result int
}

func (e *Equation) numberSolutions(withSecretElephant bool) int {
	tmp := append([]int(nil), e.terms...)
	sums := []int{tmp[0]}
	tmp = tmp[1:]
	var elem int

	for len(tmp) > 0 {

		elem, tmp = tmp[0], tmp[1:]
		newSums := []int{}
		for _, v := range sums {
			newSums = append(newSums, v+elem, v*elem)
			if withSecretElephant {
				concat, _ := strconv.Atoi(strconv.Itoa(v) + strconv.Itoa(elem))
				newSums = append(newSums, concat)
			}
		}
		sums = newSums
	}

	numSolutions := 0
	for _, v := range sums {
		if v == e.result {
			numSolutions++
		}
	}
	return numSolutions
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day7/input.txt")
	input := strings.Split(string(bytes), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func part1(input []string) int {
	total := 0
	for _, v := range input {
		eq := parseEquation(v)
		if eq.numberSolutions(false) > 0 {
			total += eq.result
		}
	}
	return total
}

func part2(input []string) int {
	total := 0
	for _, v := range input {
		eq := parseEquation(v)
		if eq.numberSolutions(true) > 0 {
			total += eq.result
		}
	}
	return total
}

func parseEquation(v string) Equation {
	parts := strings.Split(v, ":")
	result, _ := strconv.Atoi(parts[0])
	vals := strings.Fields(parts[1])
	ints := make([]int, len(vals))
	for i, val := range vals {
		ints[i], _ = strconv.Atoi(val)
	}
	return Equation{ints, result}
}
