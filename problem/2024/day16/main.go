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
	queue := []Head{{[]Point{start}, 0}}

	for len(queue) > 0 {
		path := queue[len(queue)-1].path
		currentScore := queue[len(queue)-1].score
		queue = queue[:len(queue)-1]
		current := path[len(path)-1]

		if current == end {
			if currentScore < bestScore {
				bestScore = currentScore
				fmt.Println(bestScore)
			}
			continue
		}

		if currentScore > bestScore {
			continue
		}

		candidates := getAdjacent(current, points)
		for _, next := range candidates {
			if !contains(path, next) {
				newPath := make([]Point, len(path)+1)
				copy(newPath, path)
				newPath[len(path)] = next
				queue = append(queue, Head{newPath, currentScore + score(path, next)})
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
	res := make([]Point, 0)
	for _, diff := range candidates {
		point := current.Add(diff)
		if points[point] {
			res = append(res, point)
		}
	}
	return res
}

//359296 high
//357288 high
