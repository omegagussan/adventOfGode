package common

import "strconv"

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func AbsDiffUint(x, y uint) uint {
	if x < y {
		return y - x
	}
	return x - y
}

func ToInt(input string) int {
	i, _ := strconv.ParseInt(input, 10, 64)
	return int(i)
}

const MaxUint = ^uint(0)
const MaxInt = int(MaxUint >> 1)
