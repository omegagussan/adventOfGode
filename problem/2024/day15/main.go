package main

import (
	"fmt"
	"os"
	"strings"
)

type Vector struct {
	x, y int
}

func (v Vector) add(v2 Vector) Vector {
	return Vector{v.x + v2.x, v.y + v2.y}
}
func (v Vector) scale(s int) Vector {
	return Vector{v.x * s, v.y * s}
}

func (v Vector) distance(o Vector) int {
	return (v.x-o.x)*(v.x-o.x) + (v.y-o.y)*(v.y-o.y)
}

func parseMap(input string, part2 bool) [][]string {
	half := strings.Split(input, "\n\n")
	res := make([][]string, 0, len(half[0]))
	for _, v := range strings.Split(half[0], "\n") {
		t := make([]string, 0)
		for _, c := range v {
			//If the tile is #, the new map contains ## instead.
			//If the tile is O, the new map contains [] instead.
			//If the tile is ., the new map contains .. instead.
			//If the tile is @, the new map contains @. instead.
			if part2 && c == '#' {
				t = append(t, "#")
				t = append(t, "#")
			} else if part2 && c == 'O' {
				t = append(t, "[")
				t = append(t, "]")
			} else if part2 && c == '.' {
				t = append(t, ".")
				t = append(t, ".")
			} else if part2 && c == '@' {
				t = append(t, "@")
				t = append(t, ".")
			} else {
				t = append(t, string(c))
			}
		}
		res = append(res, t)
	}
	return res
}

func parseSequence(input string) string {
	half := strings.Split(input, "\n\n")
	return strings.Replace(half[1], "\n", "", -1)
}

func sumGPS(mapz [][]string) int {
	sum := 0
	for i, row := range mapz {
		for j, cell := range row {
			if cell == "O" || cell == "[" {
				sum += 100*i + j
			}
		}
	}
	return sum
}

func part1(mapz [][]string, seq string) int {
	printMap(mapz)
	for _, a := range seq {
		curr := getRobotCoordinates(mapz)
		dir := nextStep(string(a))
		fmt.Println("dir: ", string(a))
		next := curr.add(dir)
		if getValue(next, mapz) == "." {
			mapz = swap(mapz, &curr, &next)
			continue
		} else if getValue(next, mapz) == "#" {
			continue
		}
		for getValue(next, mapz) == "O" {
			next = next.add(dir)
		}
		if getValue(next, mapz) == "#" {
			continue
		}
		backDir := dir.scale(-1)
		nextNext := next.add(backDir)
		for next != curr {
			mapz = swap(mapz, &nextNext, &next)
			nextNext, next = next.add(backDir), nextNext
		}
		printMap(mapz)
		fmt.Println()
	}
	return sumGPS(mapz)
}

func part2(mapz [][]string, seq string) int {
	printMap(mapz)
Outer:
	for _, a := range seq {
		curr := getRobotCoordinates(mapz)
		dir := nextStep(string(a))
		next := curr.add(dir)
		if getValue(next, mapz) == "." {
			mapz = swap(mapz, &curr, &next)
			continue
		} else if getValue(next, mapz) == "#" {
			continue
		}
		if string(a) == "<" || string(a) == ">" {
			for getValue(next, mapz) == "[" || getValue(next, mapz) == "]" {
				next = next.add(dir.scale(2))
			}
			if getValue(next, mapz) == "#" {
				continue
			}
			backDir := dir.scale(-1)
			nextNext := next.add(backDir)
			for next != curr {
				mapz = swap(mapz, &nextNext, &next)
				nextNext, next = next.add(backDir), nextNext
			}
			continue
		}
		//key is column
		backshifts := make(map[int]int)
		heads := make([]Vector, 0)
		heads = append(heads, next)
		for len(heads) > 0 {
			next = heads[0]
			heads = heads[1:]
			if getValue(next, mapz) == "#" {
				continue Outer
			} else if getValue(next, mapz) == "]" {
				heads = append(heads, next.add(Vector{-1, 0}).add(dir))
				heads = append(heads, next.add(dir))
			} else if getValue(next, mapz) == "[" {
				heads = append(heads, next.add(Vector{1, 0}).add(dir))
				heads = append(heads, next.add(dir))
			} else {
				//check if already exists in back-shifts.
				old := backshifts[next.x]
				if old-curr.y < next.y-curr.y {
					backshifts[next.x] = next.y
				}
			}
		}
		printMap(mapz)
		backDir := dir.scale(-1)
		for x, y := range backshifts {
			next = Vector{x, y}
			nextNext := next.add(backDir)
			diff := abs(curr.x - x)
			for next.y != curr.add(dir.scale(1+diff)).y {
				mapz = swap(mapz, &nextNext, &next)
				nextNext, next = next.add(backDir), nextNext
			}
		}
		tmp := curr.add(dir)
		mapz = swap(mapz, &curr, &tmp)
	}
	return sumGPS(mapz)
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}
func swap(mapz [][]string, a, b *Vector) [][]string {
	tmp := mapz[b.y][b.x]
	mapz[b.y][b.x] = mapz[a.y][a.x]
	mapz[a.y][a.x] = tmp
	return mapz
}

func nextStep(dir string) Vector {
	switch dir {
	case "^":
		return Vector{0, -1}
	case "v":
		return Vector{0, 1}
	case ">":
		return Vector{1, 0}
	case "<":
		return Vector{-1, 0}
	}
	panic("invalid direction")
}

func getValue(v Vector, mapz [][]string) string {
	if v.y < 0 || v.y >= len(mapz) || v.x < 0 || v.x >= len(mapz[v.y]) {
		panic("end of map!")
	}
	return mapz[v.y][v.x]
}

// x, y
func getRobotCoordinates(mapz [][]string) Vector {
	for y, row := range mapz {
		for x, cell := range row {
			if cell == "@" {
				return Vector{x, y}
			}
		}
	}
	panic("robot not found!")
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day15/input.txt")
	input := string(bytes)
	//mapz := parseMap(input, false)
	mapz2 := parseMap(input, true)
	seq := parseSequence(input)
	//fmt.Println(part1(mapz, seq))
	fmt.Println(part2(mapz2, seq))
}

func printMap(mapz [][]string) {
	for _, row := range mapz {
		fmt.Println(row)
	}
}
