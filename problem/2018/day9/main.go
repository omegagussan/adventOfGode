package main

import (
	"adventOfGode/common"
	"slices"
)

func main() {
	//println("highscore:", simulate(9, 32))
	//println("highscore:", simulate(10, 1618))
	//println("highscore:", simulate(13, 7999))
	//println("highscore:", simulate(30, 5807))
	//println("highscore:", simulate(425, 70848))
	println("highscore:", simulate(425, 70848*100))

}

func simulate(nrUsers int, marbleCuttoff int) int {
	state := []int{0, 2, 1, 3}
	curr := len(state) - 1
	score := make(map[int]int)
	for next := 4; next < marbleCuttoff; next++ {
		user := next % nrUsers
		if next%23 == 0 {
			tmp := next
			removeIdx := (curr - 7 + len(state)) % len(state)
			tmp += state[removeIdx]
			state = slices.Concat(state[:removeIdx], state[removeIdx+1:])
			curr = removeIdx

			score[user] += tmp

		} else {
			insertIdx := (curr + 2) % len(state)
			state = slices.Concat(state[:insertIdx], []int{next}, state[insertIdx:])
			curr = insertIdx
		}
	}
	return common.MapMax(score)
}
