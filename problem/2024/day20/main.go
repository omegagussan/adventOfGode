package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y, x int
}

type Step struct {
	point Point
	steps []Point
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day20/sample.txt")
	walls, start, end := parse(string(bytes))
	baseline := findShortestPath(walls, start, end)
	fmt.Println(start, end)
	fmt.Println(len(baseline))
	fmt.Println(part1(walls, start, end, len(baseline), 40))
	//add start to the start of baseline
	baseline = append([]Point{start}, baseline...)
	fmt.Println(part2(baseline, 72))
}

func part1(walls map[Point]bool, start, end Point, baseline, savings int) int {
	cheats := 0
	for wall, _ := range walls {
		walls[wall] = false
		candidate := len(findShortestPath(walls, start, end))
		walls[wall] = true
		if candidate > 0 && baseline-candidate >= savings {
			cheats++
		}
	}
	return cheats
}

func part2(baseline []Point, savings int) int {
	cheats := 0

	baselineWalls := make(map[Point]bool)
	for _, p := range baseline {
		baselineWalls[p] = true
	}
	for i, src := range baseline {
		desiredLookForward := i + savings - 1
		if desiredLookForward >= len(baseline) {
			break
		}

		for j, dst := range baseline[desiredLookForward:] {
			baselineWalls[dst] = false
			baselineWalls[src] = false
			cheatz := findAllPaths(baselineWalls, src, dst, savings-len(baseline[:i])-len(baseline[desiredLookForward+j+1:]))
			for _, cheat := range cheatz {
				// add src to the start of the cheat
				if len(cheat) > 0 && len(cheat) <= 22 {
					newPathLength := len(baseline[:i]) + len(cheat) + len(baseline[desiredLookForward+j+1:])
					newSaving := len(baseline) - newPathLength
					if newSaving >= savings {
						fmt.Println(baseline[:i], "|", cheat, "|", baseline[desiredLookForward+j+1:])
						fmt.Println(newSaving)
						cheats++
					}
				}
			}
			baselineWalls[dst] = true
			baselineWalls[src] = true
		}
	}
	return cheats
}

func neighbors(p Point) []Point {
	return []Point{
		{p.y - 1, p.x},
		{p.y + 1, p.x},
		{p.y, p.x - 1},
		{p.y, p.x + 1},
	}
}

func findShortestPath(walls map[Point]bool, start, end Point) []Point {
	// BFS
	queue := []Step{{start, []Point{}}}
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
			tmp := make([]Point, len(current.steps)+1)
			copy(tmp, current.steps)
			tmp[len(tmp)-1] = neighbor
			queue = append(queue, Step{neighbor, tmp})
		}
	}
	return []Point{}
}

func findAllPaths(walls map[Point]bool, start, end Point, earlyStop int) [][]Point {
	paths := make([][]Point, 0)
	queue := []Step{{start, []Point{start}}}
	visited := make(map[Point]bool) // Moved outside the loop

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if current.point == end {
			paths = append(paths, current.steps)
			continue
		}

		for _, neighbor := range neighbors(current.point) {
			if len(current.steps) > earlyStop {
				continue
			}

			if walls[neighbor] || visited[neighbor] {
				continue
			}

			visited[neighbor] = true // Mark as visited

			tmp := make([]Point, len(current.steps)+1)
			copy(tmp, current.steps)
			tmp[len(tmp)-1] = neighbor
			queue = append(queue, Step{neighbor, tmp})
		}
	}
	return paths
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
