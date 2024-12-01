package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day1/input.txt")
	input := string(bytes)
	println(part1(input))
	println(part2(input))
}

func part1(input string) int64 {
	sArr := strings.Split(input, "\n")
	var res int64
	for _, s := range sArr {
		i, _ := strconv.ParseInt(s, 10, 64)
		res += i
	}
	return res
}

func part2(input string) int64 {
	sArr := strings.Split(input, "\n")
	var res int64
	d := make(map[int64]bool)
	for p := 0; ; p++ {
		i, _ := strconv.ParseInt(sArr[p%len(sArr)], 10, 64)
		res += i
		if d[res] {
			break
		}
		d[res] = true
	}
	return res
}
