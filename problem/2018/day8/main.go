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
	total, _, part2 := parse(iArr, 0, make([]int, 0))
	println("part1:", total)
	println("part2:", part2)
}

func parse(data []int, total int, scores []int) (int, []int, int) {
	var nrChildren = data[0]
	var nrMetadata = data[1]
	data = data[2:]

	for _ = range nrChildren {
		var s = 0
		total, data, s = parse(data, total, make([]int, 0))
		scores = append(scores, s)
	}
	total += common.Sum(data[:nrMetadata])

	if nrChildren == 0 {
		return total, data[nrMetadata:], common.Sum(data[:nrMetadata])
	}
	var tmp = common.Map(data[:nrMetadata], func(i int) int {
		if i == 0 || i > len(scores) {
			return 0
		}
		return scores[i-1]
	})
	return total, data[nrMetadata:], common.Sum(tmp)
}

func printStringArr(iArr []int) {
	var sArr = common.Map(iArr, func(i int) string {
		return strconv.Itoa(i)
	})
	println(strings.Join(sArr, " "))
}
