package main

import (
	"os"
	"strconv"
	"strings"
)

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2017/day1/" + "input.txt")
	input := string(bytes)
	println(input)
	var stringArr = strings.Split(input, "")
	intArr := make([]int, len(stringArr))
	for i, v := range stringArr {
		intVal, _ := strconv.Atoi(v)
		intArr[i] = intVal
	}
	println("part 1: ", getSum(intArr))
	println("part 2: ", getSumPart2(intArr))
}

func getSum(arr []int) int {
	sum := 0
	for i, v := range arr {
		if i == len(arr)-1 {
			if v == arr[0] {
				sum += v
			}
		} else if v == arr[i+1] {
			sum += v
		}
	}
	return sum
}

func getSumPart2(arr []int) int {
	sum := 0
	var half = len(arr) / 2
	for i, v := range arr {
		if v == arr[(i+half)%len(arr)] {
			sum += v
		}
	}
	return sum
}
