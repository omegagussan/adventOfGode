package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`.*: X.(\d+), Y.(\d+)`)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day13/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

type Vector struct {
	x int
	y int
}

func (v Vector) Scale(i int) Vector {
	return Vector{v.x * i, v.y * i}
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

type Machine struct {
	buttons []Vector
	target  Vector
}

func (m Machine) Solve() int {
	solutions := make([]Vector, 0)
	for a := 0; a < 1000; a++ {
		if m.buttons[0].Scale(a).x > m.target.x || m.buttons[0].Scale(a).y > m.target.y {
			break
		}
		for b := 0; b < 1000; b++ {
			t := m.buttons[0].Scale(a).Add(m.buttons[1].Scale(b))
			if t == m.target {
				solutions = append(solutions, Vector{a, b})
			} else if t.x > m.target.x || t.y > m.target.y {
				break
			}
		}
	}
	if len(solutions) == 0 {
		return 0
	}
	lowestCost := int(^uint(0) >> 1)
	for _, s := range solutions {
		cost := 3*s.x + s.y
		if cost < lowestCost {
			lowestCost = cost
		}
	}
	return lowestCost
}

func part1(input string) int {
	costs := 0
	for _, ms := range strings.Split(input, "\n\n") {
		m := parseMachine(ms)
		costs += m.Solve()
	}
	return costs
}

func parseMachine(s string) Machine {
	matches := r.FindAllStringSubmatch(s, -1)

	v := make([]Vector, 0)
	for _, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		if x < 1 && y < 1 {
			continue
		}
		v = append(v, Vector{x, y})
	}
	return Machine{buttons: v[:2], target: v[2]}
}
