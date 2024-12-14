package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day14/sample.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

type Vector struct {
	x int
	y int
}

type Robot struct {
	pos, vel Vector
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

func (v Vector) Scale(i int) Vector {
	return Vector{v.x * i, v.y * i}
}

func part1(input string) int {
	robots := make([]Robot, 0)
	for _, s := range strings.Split(input, "\n") {
		robots = append(robots, parseRobots(s))
	}
	for j := 0; j < len(robots); j++ {
		r := robots[j]
		for i := 0; i < 100; i++ {
			newPos := r.pos.Add(r.vel)
			wrappedPos := wrapAround(newPos, 11, 7)
			r.pos = wrappedPos
		}
		robots[j] = r
	}
	return Score(robots, 11, 7)
}

func wrapAround(vector Vector, maxX int, maxY int) Vector {
	vector.x = (vector.x + maxX) % maxX
	vector.y = (vector.y + maxY) % maxY
	return vector
}

func Score(robots []Robot, maxX int, maxY int) int {
	quadrantCount := make([]int, 4)
	for _, r := range robots {
		quadrant := getQuadrant(r.pos, maxX, maxY)
		if quadrant > -1 {
			quadrantCount[quadrant]++
		}
	}
	score := 1
	for _, i := range quadrantCount {
		score *= i
	}
	return score
}

func getQuadrant(vector Vector, maxX int, maxY int) int {
	x := vector.x
	y := vector.y
	if x < maxX/2 && y < maxY/2 {
		return 0
	} else if x > maxX/2 && y < maxY/2 {
		return 1
	} else if x < maxX/2 && y > maxY/2 {
		return 2
	} else if x > maxX/2 && y > maxY/2 {
		return 3
	}
	return -1 // should never reach here
}
func parseRobots(s string) Robot {
	matches := r.FindAllStringSubmatch(s, -1)

	x, _ := strconv.Atoi(matches[0][1])
	y, _ := strconv.Atoi(matches[0][2])
	v1, _ := strconv.Atoi(matches[0][3])
	v2, _ := strconv.Atoi(matches[0][4])

	return Robot{Vector{x, y}, Vector{v1, v2}}
}
