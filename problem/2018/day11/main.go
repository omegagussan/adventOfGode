package main

func makeBoard(serialNr int, xMax int, yMax int) Board {
	var board = make(map[int]map[int]int)
	for x := 1; x <= xMax; x++ {
		board[x] = make(map[int]int)
		for y := 1; y <= yMax; y++ {
			board[x][y] = powerLevel(serialNr, x, y)
		}
	}
	return Board{xMax, yMax, board}
}

// x,y,power
func findMaxSquare(board Board) (int, int, int) {
	maxPower := 0
	maxX := 0
	maxY := 0
	for x := 1; x <= board.xMax-3; x++ {
		for y := 1; y <= board.yMax-3; y++ {
			pw := squarePower(board, x, y)
			if pw > maxPower {
				maxPower = pw
				maxX = x
				maxY = y
			}
		}
	}
	return maxX, maxY, maxPower
}

func squarePower(board Board, x int, y int) int {
	power := 0
	for i := x; i < x+3; i++ {
		for j := y; j < y+3; j++ {
			power += board.board[i][j]
		}
	}
	return power
}

func powerLevel(serialNr int, x int, y int) int {
	rackId := x + 10
	power := rackId * y
	power += serialNr
	power *= rackId
	power = (power / 100) % 10
	power -= 5
	return power
}

type Board struct {
	xMax  int
	yMax  int
	board map[int]map[int]int
}

func main() {
	board := makeBoard(9306, 300, 300)
	x, y, power := findMaxSquare(board)
	println("part1:", x, y, power)
}
