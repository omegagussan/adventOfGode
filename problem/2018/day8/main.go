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
	bytes, _ := os.ReadFile(dir + "/problem/2018/day8/" + "sample.txt")
	var input = string(bytes)
	var sArr = strings.Split(input, " ")
	var iArr = common.Map(sArr, func(s string) int {
		i, _ := strconv.Atoi(s)
		return i
	})
	var _, _ = parse(iArr)
}

func parse(iArr []int) (Node, int) {
	//parse node
	var i = 0
	var nrChildren = iArr[i]
	var nrMetadata = iArr[i+1]
	var n = Node{make([]Node, nrChildren), make([]int, nrMetadata)}
	i += 2
	if nrChildren == 0 {
		n.metadata = iArr[i : i+nrMetadata]
	} else {
		var size = 0
		for _ = range nrChildren {
			var latestChild, lastIndex = parse(iArr[i+size:])
			size += lastIndex
			n.children = append(n.children, latestChild)
		}
		i += size
		printStringArr(iArr)
		println("nrChildren", nrChildren)
		println("nrMetadata", nrMetadata)
		println(i)
		println(" ")
		n.metadata = iArr[i : i+nrMetadata]
	}
	return n, i
}

func printStringArr(iArr []int) {
	var sArr = common.Map(iArr, func(i int) string {
		return strconv.Itoa(i)
	})
	println(strings.Join(sArr, " "))
}
