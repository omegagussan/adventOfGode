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

func (e *Equation) numberSolutions() int {
	tmp := make([]int, len(e.terms))
	copy(tmp, e.terms)

	sums := make([]int, 0)
	//pop
	elem, tmp := tmp[0], tmp[1:]
	sums = append(sums, elem)
	for len(tmp) > 0 {
		newSums := make([]int, 0)
		//pop
		elem, tmp = tmp[0], tmp[1:]

		for v := range sums {
			newSums = append(newSums, sums[v]+elem)
			newSums = append(newSums, sums[v]*elem)
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
}

func part1(input []string) int {
	total := 0
	for _, v := range input {
		eq := parseEquation(v)
		if eq.numberSolutions() > 0 {
			total += eq.result
		}
	}
	return total
}

func parseEquation(v string) Equation {
	parts := strings.Split(v, ":")
	target := parts[0]
	result, _ := strconv.Atoi(target)
	vals := strings.Split(strings.Join(parts[1:], ""), " ")
	ints := make([]int, 0)
	for _, val := range vals {
		atoi, _ := strconv.Atoi(val)
		ints = append(ints, atoi)
	}
	eq := Equation{ints, result}
	return eq
}
