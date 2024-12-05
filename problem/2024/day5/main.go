package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day5/input.txt")
	input := string(bytes)
	fmt.Println(part1(input))
}

func part1(input string) int {
	split := strings.Split(input, "\n\n")
	rules := strings.Split(split[0], "\n")
	update := strings.Split(split[1], "\n")
	r := make(map[int][]int)
	for _, rule := range rules {
		s := strings.Split(rule, "|")
		key, _ := strconv.Atoi(s[0])
		val, _ := strconv.Atoi(s[1])
		r[key] = append(r[key], val)
	}

	total := 0
	for _, v := range update {
		valuesForUpdate := getValuesToUpdate(v)
		correctV := filterCorrectlySortedUpdates(valuesForUpdate, r)
		if len(correctV) == len(valuesForUpdate) {
			total += getValueByMiddleIndex(correctV)
		}
	}

	return total
}

func filterCorrectlySortedUpdates(valuesForUpdate []int, r map[int][]int) []int {
	correctV := make([]int, 0)
	for _, v := range valuesForUpdate {
		rules := r[v]
		if containsNone(correctV, rules) {
			correctV = append(correctV, v)
		}
	}
	return correctV
}

func getValuesToUpdate(v string) []int {
	updateV := make([]int, 0)
	for _, v := range strings.Split(v, ",") {
		val, _ := strconv.Atoi(v)
		updateV = append(updateV, val)
	}
	return updateV
}

func contains(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func containsNone(arr []int, vals []int) bool {
	for _, v := range arr {
		for _, val := range vals {
			if v == val {
				return false
			}
		}
	}
	return true
}

func getValueByMiddleIndex(arr []int) int {
	return arr[len(arr)/2]
}
