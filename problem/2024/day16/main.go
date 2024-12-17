package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
}

type Head struct {
	path  []Point
	score int
}

var candidates = []Point{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

func (p Point) Add(z Point) Point {
	return Point{p.x + z.x, p.y + z.y}
}

func (p Point) Sub(z Point) Point {
	return Point{p.x - z.x, p.y - z.y}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day16/sample.txt")
	input := string(bytes)
	fmt.Println(part(input, false))
	fmt.Println(part(input, true))
}

func part(input string, part2 bool) int {
	lines := strings.Split(input, "\n")
	points := make(map[Point]bool)
	var start, end Point

	for y, line := range lines {
		for x, curr := range line {
			switch curr {
			case 'S':
				start = Point{x, y}
			case 'E':
				end = Point{x, y}
				points[end] = true
			case '.':
				points[Point{x, y}] = true
			}
		}
	}

	bestScore := make(map[Point]int)
	queue := []Head{{[]Point{start}, 0}}
	var bestPaths []Head

	for len(queue) > 0 {
		head := queue[len(queue)-1]
		path := head.path
		currentScore := head.score
		queue = queue[:len(queue)-1]
		current := path[len(path)-1]

		if i, ok := bestScore[current]; !ok || i > currentScore {
			bestScore[current] = currentScore
			if current == end {
				bestPaths = []Head{head}
			}
		} else if currentScore-i > 1001 {
			continue
		}

		if current == end {
			if bestScore[current] == currentScore {
				bestPaths = append(bestPaths, head)
			}
			continue
		}

		for _, next := range getAdjacent(current, points) {
			if !contains(path, next) {
				newPath := make([]Point, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = next
				queue = append(queue, Head{newPath, currentScore + score(path, next)})
			}
		}
	}

	if part2 {
		set := make(map[Point]bool)
		for _, h := range bestPaths {
			for _, p := range h.path {
				set[p] = true
			}
		}
		return len(set)
	}
	return bestScore[end]
}

func contains(path []Point, p Point) bool {
	for _, px := range path {
		if p == px {
			return true
		}
	}
	return false
}

func score(p []Point, point Point) int {
	direction := Point{1, 0}
	if len(p) > 1 {
		direction = p[len(p)-1].Sub(p[len(p)-2])
	}
	if p[len(p)-1].Add(direction) == point {
		return 1
	}
	return 1001
}

func getAdjacent(current Point, points map[Point]bool) []Point {
	var res []Point
	for _, diff := range candidates {
		point := current.Add(diff)
		if points[point] {
			res = append(res, point)
		}
	}
	return res
}
