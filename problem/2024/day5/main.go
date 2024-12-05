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
	//fmt.Println(part1(input))
	fmt.Println(part2(input))

}

func part2(input string) int {
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
	for i, v := range update {
		fmt.Println(i)
		valuesForUpdate := getValuesToUpdate(v)
		correctV := optionallySortUpdates(valuesForUpdate, r)
		if len(correctV) == len(valuesForUpdate) {
			total += getValueByMiddleIndex(correctV)
		}
	}

	return total
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

func isCorrectlySortedUpdates(valuesForUpdate []int, r map[int][]int) bool {
	correctV := make([]int, 0)
	for _, v := range valuesForUpdate {
		rules := r[v]
		if containsNone(correctV, rules) {
			correctV = append(correctV, v)
		} else {
			return false
		}
	}
	return true
}
func optionallySortUpdates(valuesForUpdate []int, r map[int][]int) []int {
	if !isCorrectlySortedUpdates(valuesForUpdate, r) {
		return sortUpdates(valuesForUpdate, r)
	}
	return []int{}
}

func sortUpdates(remaining []int, r map[int][]int) []int {
	permutations := generateAllPermutations(remaining)
	for v := range permutations {
		if isCorrectlySortedUpdates(v, r) {
			return v
		}
	}
	return []int{}
}

func allPermutations(arr []int) [][]int {
	var res [][]int
	var helper func([]int, []bool, []int)

	helper = func(current []int, used []bool, arr []int) {
		if len(current) == len(arr) {
			temp := make([]int, len(current))
			copy(temp, current)
			res = append(res, temp)
			return
		}

		for i := 0; i < len(arr); i++ {
			if !used[i] {
				used[i] = true
				current = append(current, arr[i])
				helper(current, used, arr)
				current = current[:len(current)-1]
				used[i] = false
			}
		}
	}

	used := make([]bool, len(arr))
	helper([]int{}, used, arr)
	return res
}

func generateAllPermutations(arr []int) <-chan []int {
	ch := make(chan []int)
	go func() {
		defer close(ch)
		var helper func([]int, []bool, []int)

		helper = func(current []int, used []bool, arr []int) {
			if len(current) == len(arr) {
				temp := make([]int, len(current))
				copy(temp, current)
				ch <- temp
				return
			}

			for i := 0; i < len(arr); i++ {
				if !used[i] {
					used[i] = true
					current = append(current, arr[i])
					helper(current, used, arr)
					current = current[:len(current)-1]
					used[i] = false
				}
			}
		}

		used := make([]bool, len(arr))
		helper([]int{}, used, arr)
	}()
	return ch
}

func getValuesToUpdate(v string) []int {
	updateV := make([]int, 0)
	for _, v := range strings.Split(v, ",") {
		val, _ := strconv.Atoi(v)
		updateV = append(updateV, val)
	}
	return updateV
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
