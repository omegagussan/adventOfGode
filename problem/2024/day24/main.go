package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type conn struct {
	a, b      string
	operation string
	target    string
}

func (c conn) yield(state map[string]int) map[string]int {
	a := state[c.a]
	b := state[c.b]
	switch c.operation {
	case "AND":
		state[c.target] = a & b
	case "OR":
		state[c.target] = a | b
	case "XOR":
		if a != b {
			state[c.target] = 1
		} else {
			state[c.target] = 0
		}
	default:
		panic("unknown operation")
	}
	return state
}

func (c conn) isSatisfied(state map[string]int) bool {
	_, ok := state[c.a]
	if !ok {
		return false
	}
	_, ok = state[c.b]
	if !ok {
		return false
	}
	return true
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day24/input.txt")
	state, connections := parse(string(bytes))
	fmt.Println(part1(state, connections))
}

var re = regexp.MustCompile(`(.+): ([01])$`)

func parse(input string) (map[string]int, []conn) {
	parts := strings.Split(input, "\n\n")
	state := make(map[string]int)
	for _, v := range strings.Split(parts[0], "\n") {
		m := re.FindStringSubmatch(v)
		i, _ := strconv.Atoi(m[2])
		state[m[1]] = i
	}
	conns := make([]conn, 0)
	for _, v := range strings.Split(parts[1], "\n") {
		parts := strings.Split(v, " ")
		conns = append(conns, conn{parts[0], parts[2], parts[1], parts[4]})

	}
	return state, conns
}

func part1(state map[string]int, conns []conn) int {
	fmt.Println(state)
	fmt.Println(conns)
	for len(conns) > 0 {
		curr := conns[0]
		conns = conns[1:]
		if curr.isSatisfied(state) {
			state = curr.yield(state)
		} else {
			conns = append(conns, curr)
		}
	}
	//get the value of z
	acc := ""
	curr := "z00"
	for val, ok := state[curr]; ok; val, ok = state[curr] {
		acc = strconv.Itoa(val) + acc
		currVal, _ := strconv.Atoi(curr[1:])
		nextVal := strconv.Itoa(currVal + 1)
		if currVal < 9 {
			nextVal = "0" + nextVal
		}
		curr = "z" + nextVal
	}
	//convert binary to decimal
	res, _ := strconv.ParseInt(acc, 2, 64)
	return int(res)
}
