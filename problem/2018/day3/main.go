package main

import (
	"adventOfGode/common"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type pair struct {
	x int
	y int
}

var re = regexp.MustCompile(`@ ([0-9]+),([0-9]+): ([0-9]+)x([0-9]+)$`)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day3/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
}

func intz(input string) int {
	i, _ := strconv.ParseInt(input, 10, 64)
	return int(i)
}

func part1(input string) int64 {
	var sArr = strings.Split(input, "\n")

	d := make(map[pair]int)

	for _, s := range sArr {
		m := re.FindStringSubmatch(s)
		var x, y, sx, sy = intz(m[1]), intz(m[2]), intz(m[3]), intz(m[4])
		for _, i := range common.SliceTo(sx) {
			for _, j := range common.SliceTo(sy) {
				d[pair{x: x + i, y: y + j}] += 1
			}
		}
	}

	//loop over d and count the number of cells with value > 1
	var res int64 = 0
	for _, v := range d {
		if v > 1 {
			res += 1
		}
	}

	return res
}

//110195 => correct
