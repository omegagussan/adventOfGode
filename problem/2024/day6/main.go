package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x int
	y int
}

var directions = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day6/input.txt")
	input := string(bytes)
	m := parseMap(input)
	guard := findStartGuard(m)
	fmt.Println(part1(m, guard))
}

func parseMap(input string) []string {
	return strings.Split(input, "\n")
}

func findStartGuard(input []string) Point {
	for i, v := range input {
		index := strings.Index(v, "^")
		if index != -1 {
			return Point{index, i}
		}
	}
	return Point{}
}

func part1(mapz []string, guard Point) int {
	direction := 0
	visited := make(map[Point]int)
	visited[guard] = 1
	for isInMap(mapz, guard) {
		guard, direction = move(mapz, guard, direction)
		visited[guard] = 1
		fmt.Println(guard, direction)
	}
	return sumVisited(visited) - 1
}

func isInMap(input []string, p Point) bool {
	return p.y >= 0 && p.y < len(input) && p.x >= 0 && p.x < len(input[p.y])
}

func move(input []string, p Point, d int) (Point, int) {
	tmp := addPoint(p, directions[d])
	if !isInMap(input, tmp) {
		return tmp, d
	}
	mapVal := string(input[tmp.y][tmp.x])
	if mapVal != "#" {
		p = tmp
		return p, d
	} else {
		d = (d + 1) % 4
		return p, d
	}
}

func addPoint(p1 Point, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func sumVisited(visited map[Point]int) int {
	sum := 0
	for _, v := range visited {
		sum += v
	}
	return sum
}
