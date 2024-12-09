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

func (p Entry) Size() int {
	return p.BlockFiles + p.FreeSpace
}

func (p Entry) BlockSize() int {
	return p.BlockFiles
}

func (p Entry) Dump() []int {
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

func Compress(parts []Entry) []int {
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

func Fragment(parts []Entry) []int {
	memory := make([]int, 0)
	lookup := make(map[int]Entry)

	for _, p := range parts {
		memory = append(memory, p.Dump()...)
		lookup[p.ID] = p
	}

	cursor, currentId := setup(memory, lookup)
	for currentId > -1 {
		for i := 0; i < len(memory); i++ {
			if memory[i] == -1 {
				forward := getForwardRangeBound(i, memory)
				if !(forward < cursor) {
					break
				}
				sourceBlock := lookup[memory[cursor]].BlockSize()
				targetBlockSize := forward - i
				if sourceBlock <= targetBlockSize {
					for x := range sourceBlock {
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
		if forward > len(memory)-1 {
			return len(memory) - 1

		}
	}
	return forward
}

func setup(memory []int, lookup map[int]Entry) (int, int) {
	maxx := 0
	for _, v := range lookup {
		if v.ID > maxx {
			maxx = v.ID
		}
	}
	//First occurrence
	for i, v := range memory {
		if v == maxx {
			return i, maxx
		}
	}
	return -1, maxx
}

func getCursorFromID(memory []int, ID int) int {
	for i, v := range memory {
		if v == ID {
			return i
		}
	}
	return -1
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

func parseInput(input string) []Entry {
	state := make([]Entry, 0)
	ID := 0
	for len(input) > 0 {
		bf, _ := strconv.Atoi(string(input[0]))
		if len(input) == 1 {
			P := Entry{ID, bf, 0}
			state = append(state, P)
			break
		}
		fs, _ := strconv.Atoi(string(input[1]))
		P := Entry{ID, bf, fs}
		state = append(state, P)
		ID++
		input = input[2:]
	}
	return state
}

//7461441379032
//8060478710966 high
//8057088531185 high
