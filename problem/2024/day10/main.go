package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	y, x, val int
}

var dirs = []Point{{1, 0, 0}, {0, 1, 0}, {-1, 0, 0}, {0, -1, 0}}

func (p Point) Add(p2 Point) Point {
	return Point{p.y + p2.y, p.x + p2.x, p.val}
}

func (p Point) isInMap(topo []string) bool {
	return p.y >= 0 && p.y < len(topo) && p.x >= 0 && p.x < len(topo[0])
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day10/input.txt")
	topo := strings.Split(string(bytes), "\n")
	fmt.Println(part1(topo))
	fmt.Println(part2(topo))
}

func part1(topo []string) int {
	starts := findStarts(topo)
	total := 0
	for _, start := range starts {
		total += explore(topo, start, true)
	}
	return total
}

func part2(topo []string) int {
	starts := findStarts(topo)
	total := 0
	for _, start := range starts {
		total += explore(topo, start, false)
	}
	return total
}

func findStarts(topo []string) []Point {
	var starts []Point
	for y, line := range topo {
		for x, char := range line {
			if char == '0' {
				starts = append(starts, Point{y, x, 0})
			}
		}
	}
	return starts
}

func explore(topo []string, start Point, useMap bool) int {
	heads := []Point{start}
	seen := make(map[Point]bool)
	count := 0
	for len(heads) > 0 {
		curr := heads[0]
		heads = heads[1:]
		nextVal := strconv.Itoa(curr.val + 1)
		for _, dir := range dirs {
			next := curr.Add(dir)
			if !next.isInMap(topo) {
				continue
			}
			if curr.val == 8 && string(topo[next.y][next.x]) == "9" {
				if useMap {
					seen[next] = true
				} else {
					count++
				}
			} else if string(topo[next.y][next.x]) == nextVal {
				heads = append(heads, Point{next.y, next.x, curr.val + 1})
			}
		}
	}
	if useMap {
		return len(seen)
	}
	return count
}
