package main

import (
	"adventOfGode/common"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	children []Node
	metadata []int
}

func (n Node) getSpace() int {
	var sum = 2
	for _, c := range n.children {
		sum += c.getSpace()
	}
	return sum + len(n.metadata)
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day8/" + "input.txt")
	var input = string(bytes)
	var sArr = strings.Split(input, " ")
	var iArr = common.Map(sArr, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
	total, _ := parse(iArr, 0)
	println("part1:", total)
}

func parse(data []int, total int) (int, []int) {
	var nrChildren = data[0]
	var nrMetadata = data[1]
	data = data[2:]

	for _ = range nrChildren {
		total, data = parse(data, total)
	}
	printStringArr(data[:nrMetadata])
	total += common.Sum(data[:nrMetadata])

	return total, data[nrMetadata:]
}

func printStringArr(iArr []int) {
	var sArr = common.Map(iArr, func(i int) string {
		return strconv.Itoa(i)
	})
	println(strings.Join(sArr, " "))
}
