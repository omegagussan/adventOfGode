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
	walls, start, end, free := parse(string(bytes))
	baseline := findShortestPath(walls, start, end)
	fmt.Println(part1(walls, start, end, len(baseline), 40))
	//add start to the start of baseline
	baseline = append([]Point{start}, baseline...)
	free[end] = true
	fmt.Println(part2(baseline, walls, free, 72))
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

func manhattan(p1, p2 Point) int {
	return abs(p1.y-p2.y) + abs(p1.x-p2.x)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part2(baseline []Point, walls, free map[Point]bool, savings int) int {
	cheats := 0

	for i, src := range baseline[:len(baseline)-savings+1] {
		for dst, _ := range free {
			if src == dst {
				continue
			}
			if manhattan(src, dst) > 20 {
				continue
			}
			free[dst] = false
			cheat := findShortestPath(free, src, dst)
			free[dst] = true
			if len(cheat) > 0 && len(cheat) <= 20 {
				rest := findShortestPath(walls, dst, baseline[len(baseline)-1])
				if len(rest) > 0 {
					newPathLen := len(cheat) + len(rest) + len(baseline[:i+1])
					if len(baseline)-newPathLen >= savings {
						fmt.Println(len(baseline) - newPathLen)
						cheats++
					}
				} else if dst == baseline[len(baseline)-1] {
					newPathLen := len(cheat) + len(baseline[:i+1])
					if len(baseline)-newPathLen >= savings {
						fmt.Println(len(baseline) - newPathLen)
						cheats++
					}
				}
			}
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

func parse(input string) (map[Point]bool, Point, Point, map[Point]bool) {
	walls := make(map[Point]bool)
	free := make(map[Point]bool)
	var start, end Point
	for y, line := range strings.Split(input, "\n") {
		for x, char := range line {
			if char == '#' {
				walls[Point{y, x}] = true
			} else if char == 'S' {
				start = Point{y, x}
			} else if char == 'E' {
				end = Point{y, x}
			} else {
				free[Point{y, x}] = true
			}
		}
	}
	return walls, start, end, free
}
