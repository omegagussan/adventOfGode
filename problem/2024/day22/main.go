package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type sequence [4]int

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day22/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println("Part1", part1(split))
	fmt.Println("Part2", part2(split))
}

func part1(split []string) int {
	sum := 0
	for _, s := range split {
		init, _ := strconv.Atoi(s)
		for i := 0; i < 2000; i++ {
			init = next(init)
		}
		sum += init
	}
	return sum
}

func getLastDigit(num int) int {
	str := strconv.Itoa(num)
	i, _ := strconv.Atoi(string(str[len(str)-1]))
	return i
}

func part2(split []string) int {
	mapz := make(map[sequence]int)

	for _, s := range split {
		prices, diffs := pricesAndDiffs(s, 2000)
		m := bestSequence(prices, diffs)
		for k, v := range m {
			mapz[k] += v
		}
	}

	highest := 0
	for _, sumVal := range mapz {
		if sumVal > highest {
			highest = sumVal
		}
	}
	return highest
}

func bestSequence(prices, changes []int) map[sequence]int {
	m := make(map[sequence]int)
	for i := 4; i < len(changes); i++ {
		seq := sequence{
			changes[i-3],
			changes[i-2],
			changes[i-1],
			changes[i],
		}
		//pick the first occurance
		if _, exists := m[seq]; !exists {
			m[seq] = prices[i]
		}
	}

	return m
}

func pricesAndDiffs(s string, l int) ([]int, []int) {
	prices := make([]int, l)
	diffs := make([]int, l)
	init, _ := strconv.Atoi(s)
	prices[0] = init % 10
	diffs[0] = 0
	for i := 1; i < l; i++ {
		init = next(init)
		prices[i] = init % 10
		diffs[i] = prices[i] - prices[i-1]
	}
	return prices, diffs
}

func next(secret int) int {
	secret = mix(64*secret, secret)
	secret = prune(secret)
	secret = mix(secret/32, secret)
	secret = prune(secret)
	secret = mix(2048*secret, secret)
	return prune(secret)
}

func mix(given, secret int) int {
	return given ^ secret
}

func prune(secret int) int {
	return secret % 16777216
}
