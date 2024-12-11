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
	//fmt.Println(driver(input, 25))
	fmt.Println(driver(input, 75))
}

func driver(input string, length int) int {
	state := parseInput(input)
	for i := 0; i < length; i++ {
		fmt.Println(i)
		state = simulate(state)
	}
	return len(state)
}

func simulate(state []int) []int {
	pow10Cache := make(map[int]int)
	for i := 0; i < len(state); i++ {
		v := state[i]
		length := intLength(v)
		if v == 0 {
			state[i] = 1
		} else if length%2 == 0 {
			l2 := length / 2
			state = insert(state, i+1, state[i]%getTenBaseRaisedTo(l2, pow10Cache))
			state[i] = state[i] / getTenBaseRaisedTo(l2, pow10Cache)
			i++
		} else {
			state[i] = v * 2024
		}
	}
	return state
}

func getTenBaseRaisedTo(length int, pow10Cache map[int]int) int {
	if val, exists := pow10Cache[length]; exists {
		return val
	}

	result := 1
	for i := 0; i < length; i++ {
		result *= 10
	}

	pow10Cache[length] = result
	return result
}

func intLength(i int) int {
	if i == 0 {
		return 1
	}
	count := 0
	for i != 0 {
		i /= 10
		count++
	}
	return count
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
