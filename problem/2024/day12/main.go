package main

import (
	"fmt"
	"os"
	"strings"
)

var alphabetUppercase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

var adjecent = []Point{
	{0, 1},
	{0, -1},
	{1, 0},
	{-1, 0},
}

type Point struct {
	y, x int
}

func (p Point) add(p2 Point) Point {
	return Point{p.y + p2.y, p.x + p2.x}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day12/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
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
		for i, _ := range regions {
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
			continue
		}
		points = append(points, point)
		newClusterThreshold--
	}
	return regions
}

func isAdjacentToRegion(point Point, region map[Point]bool) bool {
	for _, d := range adjecent {
		t := point.add(d)
		if region[t] {
			return true
		}
	}
	return false
}

func scoreRegion(region map[Point]bool) int {
	area := len(region)
	perimeter := 0
	for point, _ := range region {
		p := numberNonAdjacent(point, region)
		perimeter += p
	}
	fmt.Println(area, perimeter)
	return area * perimeter
}

func numberNonAdjacent(point Point, region map[Point]bool) int {
	count := 4
	for _, d := range adjecent {
		t := point.add(d)
		if region[t] {
			count--
		}
	}
	return count
}

func part1(input string) int {
	mapz := strings.Split(input, "\n")

	cost := 0
	for _, c := range alphabetUppercase {
		points := make([]Point, 0)
		for y, row := range mapz {
			for x, v := range row {
				if v == c {
					points = append(points, Point{y, x})
				}
			}
		}
		regions := parseRegions(points)
		for _, region := range regions {
			cost += scoreRegion(region)
		}
	}
	return cost
}

//1419936 too low
//1424228 too low
//1446000 too low
