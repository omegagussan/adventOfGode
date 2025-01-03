package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type State struct {
	p Point
	d int
}

var directions = []Point{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day6/input.txt")
	input := string(bytes)
	m := parseMap(input)
	visited, _ := part1(m, findStartGuard(m))
	up := uniquePose(visited)
	fmt.Println(len(up) - 1)
	fmt.Println(part2(m, findStartGuard(m), up))
}

func parseMap(input string) []string {
	return strings.Split(input, "\n")
}

func findStartGuard(input []string) Point {
	for i, v := range input {
		if index := strings.Index(v, "^"); index != -1 {
			return Point{index, i}
		}
	}
	return Point{}
}

func part1(mapz []string, guard Point) (map[State]struct{}, bool) {
	direction := 0
	visitedSet := make(map[State]struct{})
	for isInMap(mapz, guard) {
		guard, direction = move(mapz, guard, direction)
		if _, exists := visitedSet[State{guard, direction}]; exists {
			return visitedSet, true
		}
		visitedSet[State{guard, direction}] = struct{}{}
	}
	return visitedSet, false
}

func part2(mapz []string, guard Point, visited map[Point]struct{}) int {
	// remove guard from visited
	delete(visited, guard)

	infLoops := 0
	for v := range visited {
		if !isInMap(mapz, v) || mapz[v.y][v.x] == '#' {
			continue
		}
		_, b := part1(insertBarrel(copyMap(mapz), v), guard)
		if b {
			infLoops++
		}
	}
	return infLoops
}

func copyMap(mapz []string) []string {
	tmpMap := make([]string, len(mapz))
	for i, v := range mapz {
		tmpMap[i] = strings.Clone(v)
	}
	return tmpMap
}

func insertBarrel(tmpMap []string, n Point) []string {
	tmpMap[n.y] = replaceAtIndex(tmpMap[n.y], '#', n.x)
	return tmpMap
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

	//look before you leap
	if !isInMap(mapz, tmp) {
		return tmp, d
	}
	if mapz[tmp.y][tmp.x] != '#' {
		return tmp, d
	}

	return p, (d + 1) % 4
}

func addPoint(p1, p2 Point) Point {
	return Point{p1.x + p2.x, p1.y + p2.y}
}

func uniquePose(state map[State]struct{}) map[Point]struct{} {
	tmp := make(map[Point]struct{})
	for k := range state {
		tmp[k.p] = struct{}{}
	}
	return tmp
}

//1784 too high
//1587 too high
