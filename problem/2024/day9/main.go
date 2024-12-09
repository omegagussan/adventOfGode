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

func (p Part) BlockSize() int {
	return p.BlockFiles
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

func Fragment(parts []Part) []int {
	memory := make([]int, 0)
	lookup := make(map[int]Part)

	for _, p := range parts {
		memory = append(memory, p.Dump()...)
		lookup[p.ID] = p
	}

	cursor := getHeadOfBlock(memory, len(memory)-1)
	for cursor > 0 {
	forwards:
		for i := 0; i < len(memory); i++ {
			if memory[i] == -1 {
				forward := i
				for memory[forward] == -1 {
					forward++
					if forward > cursor {
						break forwards
					}
				}
				if forward > cursor {
					break
				}
				if lookup[memory[cursor]].BlockSize() <= forward-i {
					for x := range lookup[memory[cursor]].BlockSize() {
						Swap(memory, cursor+x, i+x)
					}
					cursor = getHeadOfBlock(memory, cursor)
				}
			}
		}
		cursor--
		cursor = getHeadOfBlock(memory, cursor)
	}
	return memory
}

func getHeadOfBlock(memory []int, cursor int) int {
	for memory[cursor] == -1 {
		cursor--
		if cursor <= 0 {
			return cursor
		}
	}
	// this should get us to the front of the block
	old := memory[cursor]
	for memory[cursor] == old {
		cursor--
		if cursor <= 0 {
			return cursor
		}
	}
	cursor++
	return cursor
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
	fmt.Println(part2(input))
}

func part1(input string) int {
	state := parseInput(input)
	compressed := Compress(state)
	return CheckSum(compressed)
}

func part2(input string) int {
	state := parseInput(input)
	compressed := Fragment(state)
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

//7461441379032
