package main

import (
	"adventOfGode/common"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	x, y int
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
	bytes, _ := os.ReadFile(dir + "/problem/2024/day16/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	lines := strings.Split(input, "\n")
	points := make(map[Point]bool)
	start := Point{}
	end := Point{}
	for y := range lines {
		for x := range lines[y] {
			curr := lines[y][x]
			if curr == 'S' {
				start = Point{x, y}
			} else if curr == 'E' {
				end = Point{x, y}
				points[end] = true
			} else if curr == '.' {
				points[Point{x, y}] = true
			}
		}
	}

	//initialize to int max
	bestScore := common.MaxInt
	queue := [][]Point{{start}}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		current := path[len(path)-1]

		if current == end {
			currentScore := score(path)
			if currentScore < bestScore {
				bestScore = currentScore
			}
			continue
		}

		candidates := getAdjacent(current, points)
		for _, next := range candidates {
			if !contains(path, next) {
				newPath := make([]Point, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = next
				queue = append(queue, newPath)
			}
		}
	}

	return bestScore
}

func contains(path []Point, p Point) bool {
	for _, px := range path {
		if p == px {
			return true
		}
	}
	return false
}

func score(p []Point) int {
	direction := Point{1, 0}
	score := 0
	for i := 1; i < len(p); i++ {
		//every turn costs 1000 points
		//every step costs 1
		curr := p[i]
		old := p[i-1]
		if curr == old.Add(direction) {
			score++
		} else {
			score += 1001
			direction = curr.Sub(old)
		}
	}
	return score
}

func getAdjacent(current Point, points map[Point]bool) []Point {
	res := make([]Point, 0)
	for _, diff := range candidates {
		point := current.Add(diff)
		if points[point] {
			res = append(res, point)
		}
	}
	return res
}
