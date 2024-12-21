package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day21/sample.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	sum := 0
	for _, line := range split {
		i, _ := strconv.Atoi(line[:len(line)-1])
		t := numericalToDirectional(line)
		fmt.Println(t, len(t))
		t = directionalToDirectional(t)
		fmt.Println(t, len(t))
		t = directionalToDirectional(t)
		fmt.Println(t, len(t))
		fmt.Println(" ")
		sum += i * len(t)
	}
	return sum
}

func numericalToDirectional(seq string) string {
	curr := 'A'
	output := ""
	for _, num := range seq {
		output += goToNumerical(num, curr, "") + "A"
		curr = num
	}
	return output
}

func directionalToDirectional(seq string) string {
	curr := 'A'
	output := ""
	for _, num := range seq {
		output += goToDirectional(num, curr, "") + "A"
		curr = num
	}
	return output
}

//	+---+---+
//	| ^ | A |
//
// +---+---+---+
// | < | v | > |
// +---+---+---+
func goToDirectional(target rune, curr rune, state string) string {
	if target == curr {
		return state
	}
	if curr == 'A' && target == '^' {
		return "<"
	} else if curr == 'A' && target == '>' {
		return "v"
	} else if curr == '>' && target == 'A' {
		return "^"
	} else if curr == '^' && target == 'A' {
		return ">"
	} else if curr == 'A' {
		return "<" + goToDirectional(target, '^', state)
	} else if target == 'A' {
		return goToDirectional('^', curr, state) + ">"
	} else if curr == 'v' {
		return string(target)
	} else if target == 'v' {
		return toMiddleInDirectional(curr)
	} else {
		return toMiddleInDirectional(curr) + goToDirectional(target, 'v', state)
	}
}

func toMiddleInDirectional(curr rune) string {
	if curr == '^' {
		return "v"
	} else if curr == '>' {
		return "<"
	} else if curr == '<' {
		return ">"
	}
	panic("invalid direction")
}

// +---+---+---+
// | 7 | 8 | 9 |
// +---+---+---+
// | 4 | 5 | 6 |
// +---+---+---+
// | 1 | 2 | 3 |
// +---+---+---+
//
//	| 0 | A |
//	+---+---+
func goToNumerical(target rune, curr rune, state string) string {
	if target == curr {
		return state
	}
	if curr == 'A' && target == '0' {
		return "<"
	} else if target == 'A' && curr == '0' {
		return ">"
	} else if target == 'A' && curr == '3' {
		return "v"
	} else if target == 'A' {
		return goToNumerical('3', curr, state) + "v"
	} else if curr == 'A' {
		return "^" + goToNumerical(target, '3', state)
	} else if target == '0' {
		return goToNumerical('2', curr, state) + "v"
	} else if curr == '0' {
		return "^" + goToNumerical(target, '2', state)
	} else if curr-target > 2 {
		state += "v"
		curr -= 3
	} else if target-curr > 2 {
		state += "^"
		curr += 3
	} else if curr-target < 0 {
		state += ">"
		curr++
	} else {
		state += "<"
		curr--
	}
	return goToNumerical(target, curr, state)
}
