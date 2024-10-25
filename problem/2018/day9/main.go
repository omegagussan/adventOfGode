package main

import (
	"adventOfGode/common"
	"container/ring"
)

func main() {
	println("part2:", simulate(425, 70848*100))

}

func simulate(nrUsers int, marbleCuttoff int) int {
	circle := ring.New(1)
	circle.Value = 0

	score := make(map[int]int)
	for next := 1; next < marbleCuttoff; next++ {
		user := next % nrUsers
		if next%23 == 0 {
			circle = circle.Move(-8)
			remove := circle.Unlink(1)
			score[user] += next + remove.Value.(int)
			circle = circle.Move(1)
		} else {
			circle = circle.Move(1)
			s := &ring.Ring{Value: next}
			circle.Link(s)
			circle = circle.Move(1)
		}
	}
	return common.MapMax(score)
}
