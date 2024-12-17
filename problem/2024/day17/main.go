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
	bytes, _ := os.ReadFile(dir + "/problem/2024/day17/sample.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) string {
	s, i := parse(input)
	fmt.Println(i)
	fmt.Println(s)
	output := ""
	p := 0
	for p < len(i) {
		v := i[p]
		if v == 0 {
			divisor := math.Pow(2, float64(v))
			res := int(float64(s[4]) / divisor)
			s[4] = res
		} else if v == 1 {
			//bitwise xor
			s[5] = s[5] ^ v
		} else if v == 2 {
			s[5] = s[5] % 8
		} else if v == 3 && s[4] != 0 {
			if p != v {
				p = v
			} else {
				p += 2
			}
		} else if v == 4 {
			s[5] = s[5] ^ s[6]
		} else if v == 5 {
			j := s[5] % 8
			output += strconv.Itoa(j)
		} else if v == 6 {
			s[5] = s[4]/2 ^ s[6]
		}
		p++
	}
	return output
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
