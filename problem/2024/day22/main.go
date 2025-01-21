package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day22/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println("Part1", part1(split))
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

func next(secret int) int {
	secret = mix(64*secret, secret)
	secret = prune(secret)
	r := secret / 32
	secret = mix(r, secret)
	secret = prune(secret)
	secret = mix(2048*secret, secret)
	secret = prune(secret)
	return secret
}

func mix(given, secret int) int {
	//bitwise XOR
	secret = given ^ secret
	return secret
}

func prune(secret int) int {
	//bitwise AND
	secret = secret % 16777216
	return secret
}
