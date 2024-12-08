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

func (v Vector) Scale(i int) Vector {
	return Vector{v.y * i, v.x * i}
}
func (p Point) Add(v Vector) Point {
	return Point{p.y + v.y, p.x + v.x}
}

func (p Point) Dist(point Point) Vector {
	return Vector{point.y - p.y, point.x - p.x}
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day8/input.txt")
	input := string(bytes)
	antennas = getAntennas(input)
	grid = strings.Split(input, "\n")
	//fmt.Println(antennas)
	//fmt.Println(grid)
	fmt.Println(part1(antennas, grid))
	fmt.Println(part2(antennas, grid))

}

func part1(antennas string, grid []string) int {
	antiNodes := make(map[Point]int)
	for _, v := range antennas {
		points := getAntennaPoints(grid, v)
		//fmt.Println(points)
		allPairs := allPairs(points)
		//fmt.Println(allPairs)
		for _, pz := range allPairs {
			for _, p := range findAntiPoints(pz[0], pz[1]) {
				if isInGrid(p, grid) {
					antiNodes[p]++
				}
			}
		}
	}
	return len(antiNodes)
}

func part2(antennas string, grid []string) int {
	antiNodes := make(map[Point]int)
	for _, v := range antennas {
		points := getAntennaPoints(grid, v)
		//fmt.Println(points)
		allPairs := allPairs(points)
		//fmt.Println(allPairs)
		for _, pz := range allPairs {
			for _, p := range findRepeatedAntiNode(pz[0], pz[1]) {
				antiNodes[p]++
			}
		}
	}
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

func findAntiPoints(p1 Point, p2 Point) []Point {
	points := make([]Point, 0)
	//find points twice the distance from p1 and from p2
	v := p1.Dist(p2)
	p3 := p1.Add(v.Scale(2))
	points = append(points, p3)
	p4 := p1.Add(v.Scale(-1))
	points = append(points, p4)
	return points
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findRepeatedAntiNode(p1 Point, p2 Point) []Point {
	points := make([]Point, 0)

	v := p1.Dist(p2)
	divisor := gcd(v.y, v.x)
	v = v.Scale(divisor)
	tmp := Point{p1.y, p1.x}
	for isInGrid(tmp, grid) {
		points = append(points, tmp)
		tmp = tmp.Add(v)
	}
	v = v.Scale(-1)
	tmp = Point{p1.y, p1.x}
	for isInGrid(tmp, grid) {
		points = append(points, tmp)
		tmp = tmp.Add(v)
	}
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
