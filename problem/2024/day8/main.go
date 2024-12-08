package main

import (
	"fmt"
	"os"
	"strings"
)

type Point struct {
	y, x int
}

type Vector struct {
	y, x int
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
	antennas := getAntennas(input)
	grid := strings.Split(input, "\n")
	fmt.Println(part1(antennas, grid))
	fmt.Println(part2(antennas, grid))
}

func part1(antennas string, grid []string) int {
	antiNodes := make(map[Point]int)
	for _, v := range antennas {
		points := getAntennaPoints(grid, v)
		for _, pz := range allPairs(points) {
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
		for _, pz := range allPairs(points) {
			for _, p := range findRepeatedAntiNode(pz[0], pz[1], grid) {
				antiNodes[p]++
			}
		}
	}
	return len(antiNodes)
}

func allPairs(points []Point) [][]Point {
	var pairs [][]Point
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			pairs = append(pairs, []Point{points[i], points[j]})
		}
	}
	return pairs
}

func isInGrid(p Point, grid []string) bool {
	return p.y >= 0 && p.y < len(grid) && p.x >= 0 && p.x < len(grid[0])
}

func findAntiPoints(p1, p2 Point) []Point {
	v := p1.Dist(p2)
	return []Point{p1.Add(v.Scale(2)), p1.Add(v.Scale(-1))}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func findRepeatedAntiNode(p1, p2 Point, grid []string) []Point {
	var points []Point
	v := p1.Dist(p2)
	divisor := gcd(v.y, v.x)
	v = v.Scale(divisor)
	for tmp := p1; isInGrid(tmp, grid); tmp = tmp.Add(v) {
		points = append(points, tmp)
	}
	for tmp := p1; isInGrid(tmp, grid); tmp = tmp.Add(v.Scale(-1)) {
		points = append(points, tmp)
	}
	return points
}

func getAntennaPoints(grid []string, v int32) []Point {
	var points []Point
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == byte(v) {
				points = append(points, Point{i, j})
			}
		}
	}
	return points
}

func getAntennas(input string) string {
	uniqueChars := ""
	for _, c := range input {
		if !strings.ContainsRune(uniqueChars, c) && c != '\n' && c != '.' {
			uniqueChars += string(c)
		}
	}
	return uniqueChars
}
