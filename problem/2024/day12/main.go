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

type PointCorner struct {
	point  Point
	corner int
}

func (p Point) add(p2 Point) Point {
	return Point{p.y + p2.y, p.x + p2.x}
}

func isInMap(p Point, mapz []string) bool {
	if p.y < 0 || p.y >= len(mapz) {
		return false
	}
	if p.x < 0 || p.x >= len(mapz[p.y]) {
		return false
	}
	return true
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

func scoreRegion(region map[Point]bool) int {
	area := len(region)
	perimeter := 0
	for point, _ := range region {
		p := numberNonAdjacent(point, region)
		perimeter += p
	}
	return area * perimeter
}

func scoreRegionWithDiscount(region map[Point]bool, mapz []string) int {
	area := len(region)
	perimeter := make(map[PointCorner]bool)
	for point, _ := range region {
		corner := corners(point, region, mapz)
		for _, c := range corner {
			perimeter[c] = true
		}
	}
	fmt.Println(area, perimeter)
	return area * len(perimeter)
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

func corners(p Point, region map[Point]bool, mapz []string) []PointCorner {
	adjacentPoints := getAdjacentPoints(p, region)
	if len(adjacentPoints) == 0 {
		points := make([]PointCorner, 0)
		for z := 0; z < 4; z++ {
			points = append(points, PointCorner{p, z})
		}
		return points
	} else if len(adjacentPoints) == 1 {
		if adjacentPoints[0].y == p.y {
			if adjacentPoints[0].x > p.x {
				return []PointCorner{
					{p, 1},
					{p, 2},
				}
			} else {
				return []PointCorner{
					{p, 0},
					{p, 3},
				}
			}
		} else {
			if adjacentPoints[0].y > p.y {
				return []PointCorner{
					{p, 2},
					{p, 3},
				}
			} else {
				return []PointCorner{
					{p, 0},
					{p, 1},
				}
			}
		}
	} else if len(adjacentPoints) == 2 {
		if isLine(adjacentPoints[0], adjacentPoints[1]) {
			return []PointCorner{}
		}
		return isSingleCorner(p, region, mapz)
	}

	return isInnerCorner(p, region, mapz)
}

func isSingleCorner(point Point, region map[Point]bool, mapz []string) []PointCorner {
	res := getSubSquaresWeights(point, region, mapz)
	for r := range res {
		if res[r] == 4 {
			p := make([]PointCorner, 0)
			p = append(p, PointCorner{point, r})
			return p
		}
	}
	return []PointCorner{}
}

func isInnerCorner(point Point, region map[Point]bool, mapz []string) []PointCorner {
	pointConers := make([]PointCorner, 0)
	res := getSubSquaresWeights(point, region, mapz)
	for r := range res {
		if res[r] == 3 {
			pointConers = append(pointConers, PointCorner{point, r})
		}
	}
	return pointConers
}

func getSubSquaresWeights(point Point, region map[Point]bool, mapz []string) []int {
	res := make([]int, 4)
	index := 0
	for ix := 0; ix < 2; ix++ {
		for iy := 0; iy < 2; iy++ {
			for y := -1; y < 1; y++ {
				for x := -1; x < 1; x++ {
					p := Point{point.y + iy, point.x + ix}
					t := p.add(Point{y, x})
					if isInMap(t, mapz) && region[t] {
						res[index]++
					}
				}
			}
			index++
		}
	}
	return res
}

func getAdjacentPoints(p Point, region map[Point]bool) []Point {
	adjacentPoints := make([]Point, 0)
	for _, d := range adjacent {
		t := p.add(d)
		if region[t] {
			adjacentPoints = append(adjacentPoints, t)
		}
	}
	return adjacentPoints
}

func isLine(p1 Point, p2 Point) bool {
	if p1.y == p2.y {
		return true
	}
	if p1.x == p2.x {
		return true
	}
	return false
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
