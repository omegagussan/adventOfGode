package main

import (
	"fmt"
	"os"
	"strconv"
)

type Entry struct {
	ID         int
	BlockFiles int
	FreeSpace  int
}

func (e Entry) Size() int {
	return e.BlockFiles + e.FreeSpace
}

func (e Entry) Dump() []int {
	list := make([]int, e.Size())
	for i := range list {
		if i < e.BlockFiles {
			list[i] = e.ID
		} else {
			list[i] = -1
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

func Compress(entries []Entry) []int {
	memory := make([]int, 0)
	for _, e := range entries {
		memory = append(memory, e.Dump()...)
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

func Fragment(entries []Entry) []int {
	memory := make([]int, 0)
	lookup := make(map[int]Entry)

	for _, e := range entries {
		memory = append(memory, e.Dump()...)
		lookup[e.ID] = e
	}

	cursor, currentId := getFirstId(memory)
	for currentId > -1 {
		for i := 0; i < len(memory); i++ {
			if memory[i] == -1 {
				forward := getForwardRangeBound(i, memory)
				if forward >= cursor {
					break
				}
				sourceBlock := lookup[memory[cursor]].BlockFiles
				targetBlockSize := forward - i
				if sourceBlock <= targetBlockSize {
					for x := 0; x < sourceBlock; x++ {
						Swap(memory, cursor+x, i+x)
					}
					currentId--
					cursor = getCursorFromID(memory, currentId)
				}
			}
		}
		currentId--
		cursor = getCursorFromID(memory, currentId)
	}
	return memory
}

func getForwardRangeBound(i int, memory []int) int {
	forward := i
	for memory[forward] == -1 {
		forward++
		if forward >= len(memory) {
			return len(memory) - 1
		}
	}
	return forward
}

func getFirstId(memory []int) (int, int) {
	cursor := len(memory) - 1
	for memory[cursor] == memory[len(memory)-1] {
		cursor--
	}
	cursor++
	return cursor, memory[cursor]
}

func getCursorFromID(memory []int, ID int) int {
	for i, v := range memory {
		if v == ID {
			return i
		}
	}
	return -1
}

func Swap(parts []int, a, b int) {
	parts[a], parts[b] = parts[b], parts[a]
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

func parseInput(input string) []Entry {
	state := make([]Entry, 0)
	ID := 0
	for len(input) > 0 {
		bf, _ := strconv.Atoi(string(input[0]))
		fs := 0
		if len(input) > 1 {
			fs, _ = strconv.Atoi(string(input[1]))
		}
		P := Entry{ID, bf, fs}
		state = append(state, P)
		ID++
		if len(input) > 1 {
			input = input[2:]
		} else {
			input = ""
		}
	}
	return state
}

//7461441379032
//8060478710966 high
//8057088531185 high
