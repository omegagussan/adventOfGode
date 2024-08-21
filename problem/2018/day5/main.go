package main

import (
	"adventOfGode/common"
	"os"
)

var UpperOffset = common.AbsDiffInt(int('A'), int('a'))

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day5/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
}

func part1(input string) int {
	for {
		old := input
		i := 1
		for {
			if common.AbsDiffInt(int(input[i]), int(input[i-1])) == UpperOffset {
				input = input[:i-1] + input[i+1:]
			}
			i++
			if i > len(input)-2 {
				break
			}
		}
		if old == input {
			break
		}
	}
	return len(input)
}
