package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day4/input.txt")
	input := strings.Split(string(bytes), "\n")
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func RotateClockwise(input []string) []string {
	n := len(input)
	result := make([]string, n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			result[i] += string(input[n-j-1][i])
		}
	}
	return result
}

func part2(input []string) int {
	total := 0
	for i := range input {
		for j := range input[0] {
			if IsCross(input, i, j) {
				total++
			}
		}
	}
	return total
}

func IsCross(input []string, i, j int) bool {
	if i+2 >= len(input) || j+2 >= len(input[0]) || input[i+1][j+1] != 'A' {
		return false
	}
	diagonalString := string(input[i][j]) + string(input[i+1][j+1]) + string(input[i+2][j+2])
	otherDiagonalString := string(input[i][j+2]) + string(input[i+1][j+1]) + string(input[i+2][j])
	return (diagonalString == "MAS" || diagonalString == "SAM") && (otherDiagonalString == "MAS" || otherDiagonalString == "SAM")
}

func part1(input []string) int {
	total := 0
	for i := 0; i < 2; i++ {
		for _, v := range input {
			total += strings.Count(v, "XMAS") + strings.Count(v, "SAMX")
		}
		input = RotateClockwise(input)
	}
	for i := 0; i < 2; i++ {
		for offset := 0; offset < len(input[0]); offset++ {
			total += strings.Count(getDiagonalString(input, offset), "XMAS") + strings.Count(getDiagonalString(input, offset), "SAMX")
		}
		for offset := 1; offset < len(input); offset++ {
			total += strings.Count(getOtherDiagonalString(input, offset), "XMAS") + strings.Count(getOtherDiagonalString(input, offset), "SAMX")
		}
		input = RotateClockwise(input)
	}
	return total
}

func getDiagonalString(src []string, offset int) string {
	ret := ""
	for i := 0; i < len(src) && i+offset < len(src[0]); i++ {
		ret += string(src[i][i+offset])
	}
	return ret
}

func getOtherDiagonalString(src []string, offset int) string {
	ret := ""
	for i := 0; i < len(src) && i+offset < len(src); i++ {
		ret += string(src[i+offset][i])
	}
	return ret
}
