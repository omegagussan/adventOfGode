package main

import (
	"fmt"
	"os"
	"strconv"
)

type Part struct {
	ID         int
	BlockFiles int
	FreeSpace  int
}

func (p Part) Size() int {
	return p.BlockFiles + p.FreeSpace
}

func (p Part) Dump() []int {
	list := make([]int, p.Size())
	for i := range p.Size() {
		if i < p.BlockFiles {
			list[i] = p.ID
		} else {
			list[i] = -1 //empty
		}
	}
	return list
}

func CheckSum(memory []int) int {
	sum := 0
	for i, v := range memory {
		if v != -1 {
			sum += i * v
		}
	}
	return sum
}

func Compress(parts []Part) []int {
	memory := make([]int, 0)
	for _, p := range parts {
		memory = append(memory, p.Dump()...)
	}
	cursor := len(memory) - 1
	for i := 0; i < len(memory); i++ {
		if memory[i] == -1 {
			for memory[cursor] == -1 {
				cursor--
			}
			if cursor < i {
				break
			}
			Swap(memory, cursor, i)
		}
	}
	return memory
}

func Swap(parts []int, a, b int) []int {
	parts[a], parts[b] = parts[b], parts[a]
	return parts
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day9/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	state := parseInput(input)
	compressed := Compress(state)
	return CheckSum(compressed)
}

func parseInput(input string) []Part {
	state := make([]Part, 0)
	ID := 0
	for len(input) > 0 {
		bf, _ := strconv.Atoi(string(input[0]))
		if len(input) == 1 {
			P := Part{ID, bf, 0}
			state = append(state, P)
			break
		}
		fs, _ := strconv.Atoi(string(input[1]))
		P := Part{ID, bf, fs}
		state = append(state, P)
		ID++
		input = input[2:]
	}
	return state
}
