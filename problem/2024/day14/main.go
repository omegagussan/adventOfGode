package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var r = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

type Vector struct {
	x, y int
}

type Robot struct {
	pos, vel Vector
}

func (v Vector) Add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}

func wrapAround(v Vector, maxX, maxY int) Vector {
	v.x = (v.x + maxX) % maxX
	v.y = (v.y + maxY) % maxY
	return v
}

func parseRobots(s string) Robot {
	matches := r.FindStringSubmatch(s)
	x, _ := strconv.Atoi(matches[1])
	y, _ := strconv.Atoi(matches[2])
	v1, _ := strconv.Atoi(matches[3])
	v2, _ := strconv.Atoi(matches[4])
	return Robot{Vector{x, y}, Vector{v1, v2}}
}

func getQuadrant(v Vector, maxX, maxY int) int {
	if v.x < maxX/2 && v.y < maxY/2 {
		return 0
	} else if v.x >= maxX/2 && v.y < maxY/2 {
		return 1
	} else if v.x < maxX/2 && v.y >= maxY/2 {
		return 2
	} else {
		return 3
	}
}

func Score(robots []Robot, maxX, maxY int) int {
	quadrantCount := make([]int, 4)
	for _, r := range robots {
		quadrant := getQuadrant(r.pos, maxX, maxY)
		quadrantCount[quadrant]++
	}
	score := 1
	for _, count := range quadrantCount {
		score *= count
	}
	return score
}

func part1(input string) int {
	robots := parseInput(input)
	for j := range robots {
		for i := 0; i < 100; i++ {
			robots[j].pos = wrapAround(robots[j].pos.Add(robots[j].vel), 101, 103)
		}
	}
	return Score(robots, 101, 103)
}

func part2(input string) int {
	robots := parseInput(input)
	poses := make(map[Vector]bool)
	count := 0
	for len(poses) < len(robots) {
		poses = make(map[Vector]bool)
		for j := range robots {
			robots[j].pos = wrapAround(robots[j].pos.Add(robots[j].vel), 101, 103)
			poses[robots[j].pos] = true
		}
		count++
	}
	return count
}

func parseInput(input string) []Robot {
	lines := strings.Split(input, "\n")
	robots := make([]Robot, len(lines))
	for i, line := range lines {
		robots[i] = parseRobots(line)
	}
	return robots
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day14/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
