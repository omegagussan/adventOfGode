package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Key struct {
	elem   int
	length int
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day11/input.txt")
	input := string(bytes)
	fmt.Println(driver(input, 75))
}

func driver(input string, length int) int {
	cache := make(map[Key]int)
	state := parseInput(input)
	sum := 0
	for _, v := range state {
		k := Key{v, length}
		sum += simulate(k, cache)
	}
	return sum
}

func simulate(k Key, cache map[Key]int) int {
	if cached, ok := cache[k]; ok {
		return cached
	}

	if k.length == 0 {
		return 1
	}
	if k.elem == 0 {
		kt := Key{1, k.length - 1}
		res := simulate(kt, cache)
		cache[kt] = res
		return res
	}
	ll := intLength(k.elem)
	if ll%2 != 0 {
		kt := Key{k.elem * 2024, k.length - 1}
		res := simulate(kt, cache)
		cache[kt] = res
		return res
	}
	tenPow := getTenBaseRaisedTo(ll / 2)
	firstHalfKey := Key{k.elem / tenPow, k.length - 1}
	firstHalf := simulate(firstHalfKey, cache)
	cache[firstHalfKey] = firstHalf
	secondHalfKey := Key{k.elem % tenPow, k.length - 1}
	secondHalf := simulate(secondHalfKey, cache)
	cache[secondHalfKey] = secondHalf
	return firstHalf + secondHalf
}

func getTenBaseRaisedTo(length int) int {
	result := 1
	for i := 0; i < length; i++ {
		result *= 10
	}

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
