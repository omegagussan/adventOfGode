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
	//fmt.Println(part1(input))
	fmt.Println(part2(input))
}

type Vector struct {
	x int
	y int
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

func (v Vector) Scale(i int) Vector {
	return Vector{v.x * i, v.y * i}
}

type Machine struct {
	buttons []Vector
	target  Vector
}

func (m Machine) Solve(part2 bool) int {
	d := m.buttons[0].x*m.buttons[1].y - m.buttons[1].x*m.buttons[0].y
	if d == 0 {
		return 0
	}
	A := (m.target.x*m.buttons[1].y - m.buttons[1].x*m.target.y) / d
	B := (m.buttons[0].x*m.target.y - m.target.x*m.buttons[0].y) / d
	if A < 0 || B < 0 {
		return 0
	}
	if m.buttons[0].Scale(A).Add(m.buttons[1].Scale(B)) != m.target {
		return 0
	}
	return 3*A + B
}

func part1(input string) int {
	costs := 0
	for _, ms := range strings.Split(input, "\n\n") {
		m := parseMachine(ms, false)
		costs += m.Solve(false)
	}
	return costs
}

func part2(input string) int {
	costs := 0
	for _, ms := range strings.Split(input, "\n\n") {
		m := parseMachine(ms, true)
		costs += m.Solve(true)
	}
	return costs
}

func parseMachine(s string, part2 bool) Machine {
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
	t := v[2]
	if part2 {
		t = t.Add(Vector{10000000000000, 10000000000000})
	}
	return Machine{buttons: v[:2], target: t}
}
