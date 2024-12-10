package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var dirs = []Point{
	{1, 0, 0},
	{0, 1, 0},
	{-1, 0, 0},
	{0, -1, 0},
}

type Point struct {
	y   int
	x   int
	val int
}

func (p Point) Add(p2 Point) Point {
	return Point{p.y + p2.y, p.x + p2.x, p.val}
}

func (p Point) isInMap(topo []string) bool {
	return p.y >= 0 && p.y < len(topo) && p.x >= 0 && p.x < len(topo[0])
}
func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day10/input.txt")
	input := string(bytes)
	topo := strings.Split(input, "\n")
	fmt.Println(part1(topo))
}

func contains(arr []Point, elem Point) bool {
	for _, ph := range arr {
		if ph == elem {
			return true
		}
	}
	return false
}

func part1(topo []string) int {
	starts := make([]Point, 0)
	for y, line := range topo {
		for x, char := range line {
			if char == '0' {
				p := Point{y, x, 0}
				starts = append(starts, p)
			}
		}
	}

	total := 0
	for _, start := range starts {
		heads := []Point{start}
		seen := make(map[Point]bool)
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
					seen[next] = true
				} else if string(topo[next.y][next.x]) == nextVal {
					p := Point{next.y, next.x, curr.val + 1}
					if !contains(heads, p) {
						heads = append(heads, p)
					}
				}
			}
		}
		total += len(seen)
	}
	return total
}
