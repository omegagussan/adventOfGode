package main

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	x, y int
}

func (v Vector) add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}
func (v Vector) scale(s int) Vector {
	return Vector{v.x * s, v.y * s}
}

func parseMap(input string) [][]string {
	half := strings.Split(input, "\n\n")
	res := make([][]string, len(half[0]))
	for _, v := range strings.Split(half[0], "\n") {
		t := make([]string, len(v))
		for _, c := range v {
			t = append(t, string(c))
		}
		res = append(res, t)
	}
	return res
}

func parseSequence(input string) string {
	half := strings.Split(input, "\n\n")
	return strings.Replace(half[1], "\n", "", -1)
}

func sumGPS(mapz [][]string) int {
	sum := 0
	for i, row := range mapz {
		for j, cell := range row {
			if cell == "0" {
				sum += 1000*i + j
			}
		}
	}
	return sum
}

func part1(mapz [][]string, seq string) int {
	curr := getRobotCoordinates(mapz)
	for _, a := range seq {
		dir := nextStep(string(a))
		next := curr.add(dir)
		for getValue(next, mapz) == "0" {
			curr.add(next)
		}
		dir = dir.scale(-1)
		swap(mapz, &curr, &next)
		for getValue(next, mapz) != "@" {
			next = curr.add(dir)
			swap(mapz, &curr, &next)
			curr = next
		}
	}
	return sumGPS(mapz)
}

func swap(mapz [][]string, a, b *Vector) [][]string {
	mapz[a.y][a.x], mapz[b.y][b.x] = mapz[b.y][b.x], mapz[a.y][a.x]
	return mapz
}

func nextStep(dir string) Vector {
	switch dir {
	case "^":
		return Vector{0, -1}
	case "v":
		return Vector{0, 1}
	case ">":
		return Vector{1, 0}
	case "<":
		return Vector{-1, 0}
	}
	panic("invalid direction")
}

func getValue(v Vector, mapz [][]string) string {
	if v.y < 0 || v.y >= len(mapz) || v.x < 0 || v.x >= len(mapz[v.y]) {
		panic("end of map!")
	}
	return mapz[v.y][v.x]
}

// x, y
func getRobotCoordinates(mapz [][]string) Vector {
	for y, row := range mapz {
		for x, cell := range row {
			if cell == "@" {
				return Vector{x, y}
			}
		}
	}
	panic("robot not found!")
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day15/sample.txt")
	input := string(bytes)
	mapz := parseMap(input)
	seq := parseSequence(input)
	fmt.Println(part1(mapz, seq))
}
