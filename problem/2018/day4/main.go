package main

import (
	"adventOfGode/common"
	"os"
	"regexp"
	"slices"
	"strings"
)

var numberGrab = regexp.MustCompile("[0-9]+")

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2018/day4/" + "input.txt")
	var input = string(bytes)
	println(part1(input))
	println(part2(input))

}

func dictSum(input map[int]int) int {
	var sum = 0
	for _, v := range input {
		sum += v
	}
	return sum
}

func part1(input string) int64 {
	var sArr = strings.Split(input, "\n")
	slices.Sort(sArr)

	//key: guard
	//value: map[minute]count
	d := parseSchedule(sArr)

	maxGuardIdx := getMaxGuard(d)
	maxGuard := d[maxGuardIdx]

	//find the minute the guard is most asleep
	maxMinute := getMaxMinute(maxGuard)

	return int64(maxGuardIdx * maxMinute)
}

func part2(input string) int64 {
	var sArr = strings.Split(input, "\n")
	slices.Sort(sArr)

	//key: guard
	//value: map[minute]count
	d := parseSchedule(sArr)

	maxGuard := 0
	maxFreq := 0
	for g, v := range d {
		for _, i := range v {
			if i > maxFreq {
				maxGuard = g
				maxFreq = i
			}
		}
	}

	//find the minute the guard is most asleep
	maxMinute := getMaxMinute(d[maxGuard])

	return int64(maxMinute * maxGuard)
}

func parseSchedule(sArr []string) map[int]map[int]int {
	d := make(map[int]map[int]int)

	var guard = 0
	var i = 0

	//while
	for {
		//handle guard-line
		m := numberGrab.FindAllStringSubmatch(sArr[i], -1)
		if len(m) > 5 {
			guard = common.ToInt(strings.Join(m[5], ""))
			i += 1
			continue
		}

		//handle sleep-line
		var from = getLineValue(sArr[i])

		//assumed its always followed by a wake-line
		to := getLineValue(sArr[i+1])

		for j := from; j < to; j++ {
			if d[guard] == nil {
				d[guard] = make(map[int]int)
			}
			d[guard][j] += 1
		}

		//always step two lines since we read the two at the time
		i += 2

		//make sure we have two lines left if need to count more sleep
		if i+2 > len(sArr) {
			break
		}
	}
	return d
}

func getMaxMinute(guard map[int]int) int {
	var maxMinute = 0
	var maxCount = 0
	for m, c := range guard {
		if c > maxCount {
			maxCount = c
			maxMinute = m
		}
	}
	return maxMinute
}

func getMaxGuard(d map[int]map[int]int) int {
	var maxGuard = 0
	var maxMinutes = 0
	for g, v := range d {
		c := dictSum(v)
		if c > maxMinutes {
			maxMinutes = c
			maxGuard = g
		}
	}
	return maxGuard
}

func getLineValue(line string) int {
	m := numberGrab.FindAllStringSubmatch(line, -1)
	return common.ToInt(strings.Join(m[4], ""))
}
