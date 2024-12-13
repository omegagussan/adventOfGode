package main

import (
	"fmt"
	"os"
	"strings"
)

var alphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
var adjacent = []Point{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type Point struct{ y, x int }
type Edge struct {
	value     Point
	direction int
}

func (p Point) add(p2 Point) Point { return Point{p.y + p2.y, p.x + p2.x} }

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day12/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseRegions(points []Point) []map[Point]bool {
	var regions []map[Point]bool
	newClusterThreshold := -1
Outer:
	for len(points) > 0 {
		point := points[0]
		points = points[1:]

		if len(regions) == 0 {
			regions = append(regions, map[Point]bool{point: true})
			continue
		}
		for i, region := range regions {
			if isAdjacentToRegion(point, region) {
				regions[i][point] = true
				newClusterThreshold = -1
				continue Outer
			}
		}

		if newClusterThreshold < 0 {
			newClusterThreshold = 10*len(points) + 10
		} else if newClusterThreshold == 0 {
			regions = append(regions, map[Point]bool{point: true})
			newClusterThreshold = -1
			continue
		}
		points = append(points, point)
		newClusterThreshold--
	}
	return regions
}

func isAdjacentToRegion(point Point, region map[Point]bool) bool {
	for _, d := range adjacent {
		if region[point.add(d)] {
			return true
		}
	}
	return false
}

func scoreRegionWithDiscount(region map[Point]bool, mapz []string) int {
	area, edgeCount := len(region), 0
	edges := make(map[Edge]bool)
	for point := range region {
		c, newEdges := countNewEdges(point, edges, region, mapz)
		edgeCount += c
		edges = newEdges
	}
	return area * edgeCount
}

func isInMap(p Point, mapz []string) bool {
	return p.y >= 0 && p.y < len(mapz) && p.x >= 0 && p.x < len(mapz[0])
}

func countNewEdges(p Point, edges map[Edge]bool, region map[Point]bool, mapz []string) (int, map[Edge]bool) {
	duplicates, before := 0, len(edges)
	for i, d := range adjacent {
		t := p.add(d)
		if !isInMap(t, mapz) || !region[t] {
			e := Edge{t, i}
			edges[e] = true
			for _, d2 := range adjacent {
				if edges[Edge{t.add(d2), i}] {
					duplicates++
				}
			}
		}
	}
	return len(edges) - before - duplicates, edges
}

func part1(input string) int {
	mapz := strings.Split(input, "\n")
	cost := 0
	for _, c := range alphabetUppercase {
		points := parsePoints(mapz, c)
		for _, region := range parseRegions(points) {
			cost += scoreRegion(region)
		}
	}
	return cost
}

func numberNonAdjacent(point Point, region map[Point]bool) int {
	count := 4
	for _, d := range adjacent {
		if region[point.add(d)] {
			count--
		}
	}
	return count
}

func scoreRegion(region map[Point]bool) int {
	area, perimeter := len(region), 0
	for point := range region {
		perimeter += numberNonAdjacent(point, region)
	}
	return area * perimeter
}

func part2(input string) int {
	mapz := strings.Split(input, "\n")
	cost := 0
	for _, c := range alphabetUppercase {
		points := parsePoints(mapz, c)
		for _, region := range parseRegions(points) {
			cost += scoreRegionWithDiscount(region, mapz)
		}
	}
	return cost
}

func parsePoints(mapz []string, c int32) []Point {
	var points []Point
	for y, row := range mapz {
		for x, v := range row {
			if v == c {
				points = append(points, Point{y, x})
			}
		}
	}
	return points
}
