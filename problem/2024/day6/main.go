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
	fmt.Println(part1(m, Point{guard.x, guard.y}))
	fmt.Println(part2(m, Point{guard.x, guard.y}))

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
	visited := make(map[Point][]int)
	visited[guard] = append(visited[guard], direction)
	for isInMap(mapz, guard) {
		guard, direction = move(mapz, guard, direction)
		visited[guard] = append(visited[guard], direction)

		//inf loop protection
		guardVals := visited[guard]
		if findAnyDuplicate(guardVals) != -1 {
			return -1
		}
	}
	return len(visited) - 1
}

func part2(mapz []string, guard Point) int {

	oldGuard := Point{guard.x, guard.y}
	direction := 0
	visited := make(map[Point][]int)
	visited[guard] = append(visited[guard], direction)
	for isInMap(mapz, guard) {
		guard, direction = move(mapz, guard, direction)
		visited[guard] = append(visited[guard], direction)
	}

	infLoops := make(map[Point]bool)
	for v, _ := range visited {
		if !isInMap(mapz, v) {
			continue
		}
		if mapz[v.y][v.x] == '#' {
			continue
		}
		tmpMap := copyMap(mapz)
		insertBarrel(tmpMap, v)
		if part1(tmpMap, Point{x: oldGuard.x, y: oldGuard.y}) == -1 {
			infLoops[v] = true
		}
	}
	return len(infLoops)
}

func copyMap(mapz []string) []string {
	var tmpMap []string
	for _, v := range mapz {
		tmpMap = append(tmpMap, strings.Clone(v))
	}
	return tmpMap
}

func insertBarrel(tmpMap []string, n Point) {
	tmpMap[n.y] = replaceAtIndex(tmpMap[n.y], '#', n.x)
}

func replaceAtIndex(in string, r rune, i int) string {
	out := []rune(in)
	out[i] = r
	return string(out)
}

func isInMap(input []string, p Point) bool {
	return p.y >= 0 && p.y < len(input) && p.x >= 0 && p.x < len(input[p.y])
}

func move(mapz []string, p Point, d int) (Point, int) {
	tmp := addPoint(p, directions[d])
	if !isInMap(mapz, tmp) {
		return tmp, d
	}
	mapVal := string(mapz[tmp.y][tmp.x])
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

//1587 => high

func findAnyDuplicate(input []int) int {
	seen := make(map[int]bool)
	for _, v := range input {
		if seen[v] {
			return v
		}
		seen[v] = true
	}
	return -1
}
