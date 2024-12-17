package main

import (
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`.*: (\d+)`)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day17/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) string {
	s, i := parse(input)
	output := ""
	p := 0
Outer:
	for p < len(i) {
		v := i[p]
		operand := i[p+1]
		switch v {
		case 0: // adv
			divisor := math.Pow(2, float64(getComboOperandValue(operand, s)))
			res := int(float64(s[4]) / divisor)
			s[4] = res
		case 1: // bxl
			s[5] = s[5] ^ operand
		case 2: // bst
			s[5] = getComboOperandValue(operand, s) % 8
		case 3: // jnz
			if s[4] != 0 {
				p = operand
				continue Outer
			}
		case 4: // bxc
			s[5] = s[5] ^ s[6]
		case 5: // out
			j := getComboOperandValue(operand, s) % 8
			output += strconv.Itoa(j)
		case 6: // bdv
			divisor := math.Pow(2, float64(getComboOperandValue(operand, s)))
			res := int(float64(s[4]) / divisor)
			s[5] = res
		case 7: // cdv
			divisor := math.Pow(2, float64(getComboOperandValue(operand, s)))
			res := int(float64(s[4]) / divisor)
			s[6] = res
		}
		p += 2
	}
	return output
}

func getComboOperandValue(operand int, s map[int]int) int {
	switch operand {
	case 0, 1, 2, 3:
		return operand
	case 4:
		return s[4]
	case 5:
		return s[5]
	case 6:
		return s[6]
	default:
		panic("this should not happen!")
	}
}

func parse(input string) (map[int]int, []int) {
	parts := strings.Split(input, "\n\n")

	//parse dict
	state := make(map[int]int)
	for x, l := range strings.Split(parts[0], "\n") {
		m := re.FindStringSubmatch(l)
		if len(m) == 2 {
			v, _ := strconv.Atoi(m[1])
			state[x+4] = v
		}
	}

	//parse instructions
	is := strings.Replace(parts[1], "Program: ", "", 1)
	instructions := make([]int, 0)

	for _, i := range strings.Split(is, ",") {
		v, _ := strconv.Atoi(i)
		instructions = append(instructions, v)
	}

	return state, instructions
}

//432645324
