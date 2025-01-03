package main

import (
	"adventOfGode/common"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y, x int
}

type Step struct {
	point Point
	steps int
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day20/input.txt")
	walls, start, end := parse(string(bytes))
	baseline := findShortestPath(walls, start, end)
	topLeft, bottomRight := findMinMax(walls)
	fmt.Println(baseline)
	fmt.Println(part1(walls, start, end, topLeft, bottomRight, baseline))
}

func part1(walls map[Point]bool, start, end, topLeft, bottomRight Point, baseline int) int {
	cheats := 0
	for wall, _ := range walls {
		if !isWithinBounds(wall, topLeft, bottomRight) {
			continue
		}
		//walls without current wall
		walls[wall] = false
		candidate := findShortestPath(walls, start, end)
		//fmt.Println(candidate)
		walls[wall] = true
		if candidate > -1 && baseline-candidate >= 100 {
			cheats++
		}
	}
	return cheats
}

func isWithinBounds(p, topLeft, bottomRight Point) bool {
	return p.y > topLeft.y && p.y < bottomRight.y && p.x > topLeft.x && p.x < bottomRight.x
}

func findMinMax(walls map[Point]bool) (Point, Point) {
	minY, minX := common.MaxInt, common.MaxInt
	maxY, maxX := 0, 0
	for point := range walls {
		if point.y < minY {
			minY = point.y
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.x < minX {
			minX = point.x
		}
		if point.x > maxX {
			maxX = point.x
		}
	}
	return Point{minY, minX}, Point{maxY, maxX}

}

func neighbors(p Point) []Point {
	return []Point{
		{p.y - 1, p.x},
		{p.y + 1, p.x},
		{p.y, p.x - 1},
		{p.y, p.x + 1},
	}
}

func findShortestPath(walls map[Point]bool, start, end Point) int {
	// BFS
	queue := []Step{{start, 0}}
	visited := make(map[Point]bool)
	visited[start] = true
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		if current.point == end {
			return current.steps
		}
		for _, neighbor := range neighbors(current.point) {
			if walls[neighbor] || visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			queue = append(queue, Step{neighbor, current.steps + 1})
		}
	}
	return -1
}

func parse(input string) (map[Point]bool, Point, Point) {
	walls := make(map[Point]bool)
	var start, end Point
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == '#' {
				walls[Point{y, x}] = true
			} else if char == 'S' {
				start = Point{y, x}
			} else if char == 'E' {
				end = Point{y, x}
			}
		}
	}
	return walls, start, end
}

//1310 too low
