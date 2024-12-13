package main

import (
	"fmt"
	"os"
	"strings"
)

var alphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var adjacent = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Point struct {
	y, x int
}

type Edge struct {
	value     Point
	direction int
}

func (p Point) add(p2 Point) Point {
	return Point{p.y + p2.y, p.x + p2.x}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day12/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}

func parseRegions(points []Point) []map[Point]bool {
	regions := make([]map[Point]bool, 0)
	newClusterThreshold := -1
Outer:
	for len(points) > 0 {
		// get the first point
		point := points[0]
		points = points[1:]

		if len(regions) == 0 {
			r := make(map[Point]bool)
			r[point] = true
			regions = append(regions, r)
			continue
		}
		for i := range regions {
			region := regions[i]
			if isAdjacentToRegion(point, region) {
				regions[i][point] = true
				newClusterThreshold = -1
				continue Outer
			}
		}

		if newClusterThreshold < 0 {
			newClusterThreshold = 10*len(points) + 10
		} else if newClusterThreshold == 0 {
			r := make(map[Point]bool)
			r[point] = true
			regions = append(regions, r)
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
		t := point.add(d)
		if region[t] {
			return true
		}
	}
	return false
}

func scoreRegionWithDiscount(region map[Point]bool, mapz []string) int {
	area := len(region)
	edgeCount := 0
	edges := make(map[Edge]bool)
	for point, _ := range region {
		c := 0
		c, edges = countNewEdges(point, edges, region, mapz)
		edgeCount += c
	}
	fmt.Println(area, edgeCount)
	return area * edgeCount
}

func getValueByDirection(p Point, direction int) int {
	switch direction {
	case 0:
		return p.x
	case 1:
		return p.x
	case 2:
		return p.y
	case 3:
		return p.y
	}
	panic("invalid direction")
}

func isInMap(p Point, mapz []string) bool {
	return p.y >= 0 && p.y < len(mapz) && p.x >= 0 && p.x < len(mapz[0])
}

func countNewEdges(p Point, edges map[Edge]bool, region map[Point]bool, mapz []string) (int, map[Edge]bool) {
	duplicates := 0
	before := len(edges)
	for i, d := range adjacent {
		t := p.add(d)
		if !isInMap(t, mapz) || !region[t] {
			e := Edge{t, i}
			edges[e] = true

			for _, d2 := range adjacent {
				t2 := t.add(d2)
				if edges[Edge{t2, i}] {
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
		regions := parseRegions(points)
		for _, region := range regions {
			cost += scoreRegion(region)
		}
	}
	return cost
}

func numberNonAdjacent(point Point, region map[Point]bool) int {
	count := 4
	for _, d := range adjacent {
		t := point.add(d)
		if region[t] {
			count--
		}
	}
	return count
}

func scoreRegion(region map[Point]bool) int {
	area := len(region)
	perimeter := 0
	for point, _ := range region {
		p := numberNonAdjacent(point, region)
		perimeter += p
	}
	return area * perimeter
}

func part2(input string) int {
	mapz := strings.Split(input, "\n")

	cost := 0
	for _, c := range alphabetUppercase {
		points := parsePoints(mapz, c)
		regions := parseRegions(points)
		for _, region := range regions {
			cost += scoreRegionWithDiscount(region, mapz)
		}
	}
	return cost
}

func parsePoints(mapz []string, c int32) []Point {
	points := make([]Point, 0)
	for y, row := range mapz {
		for x, v := range row {
			if v == c {
				points = append(points, Point{y, x})
			}
		}
	}
	return points
}

//1079711 too high
