package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct{ y, x int }

type Step struct {
	point Point
	steps int
}

func parsePoints(input string) []Point {
	splits := strings.Split(input, "\n")
	points := make([]Point, len(splits))
	for index, line := range splits {
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			continue
		}
		x, _ := strconv.Atoi(coords[0])
		y, _ := strconv.Atoi(coords[1])
		points[index] = Point{y, x}
	}
	return points
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day18/input.txt")
	start := Point{0, 0}
	goal := Point{70, 70}
	input := parsePoints(string(bytes))
	part1 := toMap(input, 1024)
	fmt.Println(findShortestPath(start, goal, part1))
	fmt.Println(findFirstNonSolution(start, goal, input[1024:], part1))
}

func findFirstNonSolution(start, goal Point, remaining []Point, walls map[Point]bool) Point {
	prev := Point{-1, -1}
	for _, curr := range remaining {
		if findShortestPath(start, goal, walls) == -1 {
			return prev
		}
		prev = curr
		walls[curr] = true
	}
	return prev
}

func toMap(points []Point, bytes int) map[Point]bool {
	r := make(map[Point]bool)
	for _, p := range points[:bytes] {
		r[p] = true
	}
	return r
}

func addPoint(m map[Point]bool, p Point) {
	m[p] = true
}

func getNeighbors(p Point) []Point {
	return []Point{{p.y + 1, p.x}, {p.y - 1, p.x}, {p.y, p.x + 1}, {p.y, p.x - 1}}
}

func isInRange(p Point, goal Point) bool {
	return p.y >= 0 && p.x >= 0 && p.y <= goal.y && p.x <= goal.x
}

func findShortestPath(start Point, goal Point, walls map[Point]bool) int {
	queue := []Step{{start, 0}}
	visited := make(map[Point]bool)
	visited[start] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if curr.point == goal {
			return curr.steps
		}

		for _, neighbor := range getNeighbors(curr.point) {
			if walls[neighbor] || !isInRange(neighbor, goal) || visited[neighbor] {
				continue
			}
			visited[neighbor] = true
			queue = append(queue, Step{neighbor, curr.steps + 1})
		}
	}
	return -1 // Return -1 if no path is found
}
