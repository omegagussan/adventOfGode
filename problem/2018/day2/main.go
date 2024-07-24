package main

import (
	"os"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day2/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
	part2(input)

}

var alphabet = "abcdefghijklmnopqrstuvwxyz"
var alphabetGen = strings.Split(alphabet, "")

func smudgeString(s string, idx int, target string) string {
	tmp := []rune(s)
	tmp[idx] = rune(target[0])
	return string(tmp)
}

func part2(input string) {
	var sArr = strings.Split(input, "\n")
	d := cache(sArr)

out:
	for _, s := range sArr {
		line := strings.Split(s, "")
		for i, _ := range line {
			for _, letter := range alphabetGen {
				candidate := smudgeString(s, i, letter)
				if candidate == s {
					continue
				}
				if d[candidate] {
					println(strings.Join(removeFromSlice(line, i), ""))
					break out
				}
			}
		}
	}
}

func cache(sArr []string) map[string]bool {
	d := make(map[string]bool)
	for _, s := range sArr {
		d[s] = true
	}
	return d
}

func removeFromSlice(s []string, idx int) []string {
	return append(s[:idx], s[idx+1:]...)

}
func part1(input string) int64 {
	var sArr = strings.Split(input, "\n")
	var pair int64 = 0
	var triplet int64 = 0

	for _, s := range sArr {
		pair += int64(exactlyMatches(s, 2))
		triplet += int64(exactlyMatches(s, 3))
	}
	return pair * triplet
}

func exactlyMatches(row string, matches int) int {
	d := make(map[string]int)
	for _, s := range strings.Split(row, "") {
		d[s] += 1
	}
	var numMatches = 0
	for _, v := range d {
		if v == matches {
			numMatches += 1
		}
	}
	if numMatches > 0 {
		return 1
	}
	return 0
}

//7688
//g y y g
// lsrivmotzbdxpkxnaqmuwc vmotzbdxpkxnaqmuwcychj
// lsrivmotzbdxpkxnaqmuwccjj
// lsrivmotzbdxpkxnaqmuwcchj
