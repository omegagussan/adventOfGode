package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Lists struct {
	First  []int
	Second []int
}

func newLists() *Lists {
	return &Lists{
		First:  []int{},
		Second: []int{},
	}
}

func (l *Lists) addFirst(i int) {
	l.First = append(l.First, i)
}

func (l *Lists) addSecond(i int) {
	l.Second = append(l.Second, i)
}

func (l *Lists) sort() *Lists {
	sort.Ints(l.First)
	sort.Ints(l.Second)
	return l
}

func (l *Lists) part1() int {
	l.sort()
	var total int
	for i, v := range l.First {
		abs := math.Abs(float64(v - l.Second[i]))
		total += int(abs)
	}
	return total
}

func (l *Lists) part2() int {
	var total int
	for _, v := range l.First {
		total += v * l.count(v)
	}
	return total
}

func (l *Lists) count(i int) int {
	total := 0
	for _, v := range l.Second {
		if i == v {
			total++
		}
	}
	return total
}

func main() {
	dir, _ := os.Getwd()
	bytes, _ := os.ReadFile(dir + "/problem/2024/day1/" + "input.txt")
	input := string(bytes)
	split := strings.Split(input, "\n")
	fmt.Println(part1(split))
	fmt.Println(part2(split))
}

func part1(split []string) int {
	l := parseLists(split)
	return l.part1()
}

func part2(split []string) int {
	l := parseLists(split)
	return l.part2()
}

func parseLists(split []string) *Lists {
	var l = newLists()
	for _, s := range split {
		nums := strings.Split(s, " ")
		first, _ := strconv.Atoi(nums[0])
		second, _ := strconv.Atoi(nums[len(nums)-1])
		l.addFirst(first)
		l.addSecond(second)
	}
	return l
}

//54925252
//2727176
