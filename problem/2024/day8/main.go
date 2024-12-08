package main

import (
	"fmt"
	"os"
	"strings"
)

var antennas string
var grid []string

type Point struct {
	y int
	x int
}

type Vector struct {
	y int
	x int
}

func (v Vector) Scale(f float32) Vector {
	return Vector{int(float32(v.y) * f), int(float32(v.x) * f)}
}
func (p Point) Add(v Vector) Point {
	return Point{p.y + v.y, p.x + v.x}
}

func (p Point) Dist(point Point) Vector {
	return Vector{p.y - point.y, p.x - point.x}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day8/sample.txt")
	input := string(bytes)
	antennas = getAntennas(input)
	grid = strings.Split(input, "\n")
	//fmt.Println(antennas)
	//fmt.Println(grid)
	fmt.Println(part1(antennas, grid))
}

func part1(antennas string, grid []string) int {
	antiNodes := make(map[Point]int)
	for _, v := range antennas {
		points := getAntennaPoints(grid, v)
		//fmt.Println(points)
		allPairs := allPairs(points)
		//fmt.Println(allPairs)
		for _, pz := range allPairs {
			for _, p := range getAntiPoints(pz[0], pz[1]) {
				if isInGrid(p, grid) {
					antiNodes[p]++
				}
			}
		}
	}
	fmt.Println(antiNodes)
	return len(antiNodes)
}

func allPairs(points []Point) [][]Point {
	pairs := make([][]Point, 0)
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, []Point{points[i], points[j]})
		}
	}
	return pairs
}

func isInGrid(p Point, grid []string) bool {
	if p.y < 0 || p.y >= len(grid) || p.x < 0 || p.x >= len(grid[0]) {
		return false
	}
	return true
}

func getAntiPoints(p1 Point, p2 Point) []Point {
	points := make([]Point, 0)
	v := p2.Dist(p1)
	v = v.Scale(1 / float32(3))
	p1a := p1.Add(v)
	points = append(points, p1a)
	p2a := p2.Add(v.Scale(-1))
	points = append(points, p2a)
	return points
}

func getAntennaPoints(grid []string, v int32) []Point {
	points := make([]Point, 0)
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if string(grid[i][j]) == string(v) {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func getAntennas(input string) string {
	uniqueChars := ""
	for _, c := range input {
		if !contains(uniqueChars, c) {
			uniqueChars += string(c)
		}
	}
	tmp := uniqueChars
	tmp = strings.Replace(tmp, "\n", "", -1)
	tmp = strings.Replace(tmp, ".", "", -1)
	return tmp
}

func contains(s string, c rune) bool {
	for _, char := range s {
		if char == c {
			return true
		}
	}
	return false
}
