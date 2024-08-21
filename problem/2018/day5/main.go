package main

import (
	"adventOfGode/common"
	"os"
	"strings"
)

var upperOffset = common.AbsDiffInt(int('A'), int('a'))
var alphabet = "abcdefghijklmnopqrstuvwxyz"
var alphabetGen = strings.Split(alphabet, "")

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day5/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
	println(part2(input))
}

func part2(input string) int {
	var m = len(input)
	for i, _ := range alphabetGen {
		temp := strings.NewReplacer(alphabetGen[i], "", strings.ToUpper(alphabetGen[i]), "").Replace(input)
		var length = part1(temp)
		if length < m {
			m = length
		}
	}
	return m
}

func part1(input string) int {
	for {
		old := input
		i := 1
		for {
			if common.AbsDiffInt(int(input[i]), int(input[i-1])) == upperOffset {
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
