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
	println(part2(input))
}

func toInt(input string) int {
	i, _ := strconv.ParseInt(input, 10, 64)
	return int(i)
}

func part1(input string) int64 {
	var sArr = strings.Split(input, "\n")

	d := make(map[pair]int)

	for _, s := range sArr {
		m := re.FindStringSubmatch(s)
		var x, y, sx, sy = toInt(m[1]), toInt(m[2]), toInt(m[3]), toInt(m[4])
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

func part2(input string) int64 {
	var sArr = strings.Split(input, "\n")

	//value is list of claim ids
	d := make(map[pair][]int)

	for claimId, s := range sArr {
		m := re.FindStringSubmatch(s)
		var x, y, sx, sy = toInt(m[1]), toInt(m[2]), toInt(m[3]), toInt(m[4])
		for _, i := range common.SliceTo(sx) {
			for _, j := range common.SliceTo(sy) {
				//append claimId to the list of claim ids at the cell
				d[pair{x: x + i, y: y + j}] = append(d[pair{x: x + i, y: y + j}], claimId)
			}
		}
	}

	for i, _ := range sArr {
		//check if unique in the list of candidates
		unique := true
		for _, d := range d {
			if len(d) > 1 && common.Contains(d, i) {
				unique = false
				break
			}
		}
		if unique {
			//index 1 based
			return int64(i + 1)
		}
	}
	return -1
}

//110195 => correct
