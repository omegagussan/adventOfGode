package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day1/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
	println(part2(input))

}

func part2(input string) int64 {
	//split string into integers
	var sArr = strings.Split(input, "\n")
	var res int64 = 0

	//there is no set in golang
	d := make(map[int64]bool)
	var p = 0
	for {
		s := sArr[p%len(sArr)]
		i, _ := strconv.ParseInt(s, 10, 64)
		res += i
		if d[res] {
			break
		}
		d[res] = true
		p += 1
	}
	return res
}

func part1(input string) int64 {
	//split string into integers
	var sArr = strings.Split(input, "\n")
	var res int64 = 0
	for _, s := range sArr {
		i, _ := strconv.ParseInt(s, 10, 64)
		res += i
	}
	return res
}

//990 high
//585 => correct. had wrap around.
