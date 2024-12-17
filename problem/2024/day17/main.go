package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type State struct {
	a, b, c int
}

func combo(v int, vm State) int {
	switch v {
	case 0, 1, 2, 3:
		return v
	case 4:
		return vm.a
	case 5:
		return vm.b
	case 6:
		return vm.c
	default:
		panic("invalid operand")
	}
}

func main() {
	dir, _ := os.Getwd()
	input, _ := os.ReadFile(dir + "/problem/2024/day17/input.txt")
	operations, state := parse(string(input))

	var output []string
	for p := 0; p < len(operations)-1; p += 2 {
		v := operations[p+1]
		comboValue := combo(v, state)
		switch operations[p] {
		case 0:
			state.a >>= comboValue
		case 1:
			state.b ^= v
		case 2:
			state.b = comboValue % 8
		case 3:
			if state.a != 0 {
				p = v - 2
			}
		case 4:
			state.b ^= state.c
		case 5:
			output = append(output, strconv.Itoa(comboValue%8))
		case 6:
			state.b = state.a >> comboValue
		case 7:
			state.c = state.a >> comboValue
		}
	}
	fmt.Println(strings.Join(output, ","))
}

func parse(inputStr string) ([]int, State) {
	var state State
	var operationString string
	fmt.Sscanf(inputStr, "Register A: %d\nRegister B: %d\nRegister C: %d\n\nProgram: %s\n", &state.a, &state.b, &state.c, &operationString)
	return parseOperations(operationString), state
}

func parseOperations(programStr string) []int {
	parts := strings.Split(programStr, ",")
	operations := make([]int, len(parts))
	for i, op := range parts {
		operations[i], _ = strconv.Atoi(op)
	}
	return operations
}
