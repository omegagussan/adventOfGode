package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day11/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	state := parseInput(input)
	for i := 0; i < 25; i++ {
		state = simulate(state)
	}
	return len(state)
}

func simulate(state []int) []int {
	for i := 0; i < len(state); i++ {
		v := state[i]
		stringV := strconv.Itoa(v)
		if v == 0 {
			state[i] = 1
		} else if len(stringV)%2 == 0 {
			firstHalf, _ := strconv.Atoi(stringV[:len(stringV)/2])
			SecondHalf, _ := strconv.Atoi(stringV[len(stringV)/2:])
			state[i] = firstHalf
			state = insert(state, i+1, SecondHalf)
			i++
		} else {
			state[i] = v * 2024
		}
	}
	return state
}

// 0 <= index <= len(a)
func insert(a []int, index int, value int) []int {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func parseInput(input string) []int {
	state := make([]int, 0)
	for _, s := range strings.Split(input, " ") {
		i, _ := strconv.Atoi(s)
		state = append(state, i)
	}
	return state
}
