package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day21/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	split := strings.Split(input, "\n")
	sum := 0
	for _, line := range split {
		fmt.Println(line)
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
		output += goToNumerical(num, curr) + "A"
		curr = num
	}
	return output
}

func directionalToDirectional(seq string) string {
	curr := 'A'
	output := ""
	for _, num := range seq {
		output += goToDirectional(num, curr) + "A"
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
func goToDirectional(target rune, curr rune) string {
	keypad := map[rune][2]int{
		'^': {0, 1}, 'A': {0, 2},
		'<': {1, 0}, 'v': {1, 1}, '>': {1, 2},
	}

	startPos := keypad[curr]
	endPos := keypad[target]

	dx := endPos[1] - startPos[1]
	dy := endPos[0] - startPos[0]

	var sb strings.Builder

	if startPos[0] == 0 {
		if dy > 0 {
			sb.WriteString(strings.Repeat("v", dy))
		} else if dy < 0 {
			sb.WriteString(strings.Repeat("^", -dy))
		}

		if dx > 0 {
			sb.WriteString(strings.Repeat(">", dx))
		} else if dx < 0 {
			sb.WriteString(strings.Repeat("<", -dx))
		}
	} else {
		if dx > 0 {
			sb.WriteString(strings.Repeat(">", dx))
		} else if dx < 0 {
			sb.WriteString(strings.Repeat("<", -dx))
		}

		if dy > 0 {
			sb.WriteString(strings.Repeat("v", dy))
		} else if dy < 0 {
			sb.WriteString(strings.Repeat("^", -dy))
		}
	}

	return sb.String()
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
func goToNumerical(end, start rune) string {
	keypad := map[rune][2]int{
		'7': {0, 0}, '8': {0, 1}, '9': {0, 2},
		'4': {1, 0}, '5': {1, 1}, '6': {1, 2},
		'1': {2, 0}, '2': {2, 1}, '3': {2, 2},
		'0': {3, 1}, 'A': {3, 2},
	}

	startPos := keypad[start]
	endPos := keypad[end]

	dx := endPos[1] - startPos[1]
	dy := endPos[0] - startPos[0]

	var sb strings.Builder

	if startPos[0] == 3 {
		if dy > 0 {
			sb.WriteString(strings.Repeat("v", dy))
		} else if dy < 0 {
			sb.WriteString(strings.Repeat("^", -dy))
		}

		if dx > 0 {
			sb.WriteString(strings.Repeat(">", dx))
		} else if dx < 0 {
			sb.WriteString(strings.Repeat("<", -dx))
		}
	} else {
		if dx > 0 {
			sb.WriteString(strings.Repeat(">", dx))
		} else if dx < 0 {
			sb.WriteString(strings.Repeat("<", -dx))
		}

		if dy > 0 {
			sb.WriteString(strings.Repeat("v", dy))
		} else if dy < 0 {
			sb.WriteString(strings.Repeat("^", -dy))
		}
	}

	return sb.String()
}

//258330 too high
