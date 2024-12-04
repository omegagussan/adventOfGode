package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day4/input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println(part1(split))
}

func RotateClockwise(input []string) []string {
	result := make([]string, len(input))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			result[i] += string(input[len(input)-j-1][i])
		}
	}
	return result
}

func part1(input []string) int {
	total := 0
	rotated := input
	for i := 0; i < 2; i++ {
		for _, v := range rotated {
			total += strings.Count(v, "XMAS")
			total += strings.Count(v, "SAMX")
		}
		rotated = RotateClockwise(input)
	}
	println(total)

	//search diagonally
	for i := 0; i < 2; i++ {
		for offset := 0; offset < len(input[0]); offset++ {
			diagonalString := getDiagonalString(input, offset)
			println(diagonalString)
			total += strings.Count(diagonalString, "XMAS")
			total += strings.Count(diagonalString, "SAMX")
		}
		for offset := 1; offset < len(input); offset++ {
			otherDiagonalString := getOtherDiagonalString(input, offset)
			println(otherDiagonalString)
			total += strings.Count(otherDiagonalString, "XMAS")
			total += strings.Count(otherDiagonalString, "SAMX")

		}
		input = RotateClockwise(input)
	}
	return total
}

func getDiagonalString(src []string, offset int) string {
	ret := ""
	rangez := max(len(src), len(src[0]))
	for i := 0; i < rangez; i++ {
		if i+offset >= len(src[0]) {
			break
		}
		ret += string(src[i][i+offset])
	}
	return ret
}

func getOtherDiagonalString(src []string, offset int) string {
	ret := ""
	rangez := max(len(src), len(src[0]))
	for i := 0; i < rangez; i++ {
		if i+offset >= len(src) {
			break
		}
		ret += string(src[i+offset][i])
	}
	return ret
}
