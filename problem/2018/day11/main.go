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
func findMaxSquare(board Board, size int) (int, int, int) {
	maxPower := 0
	maxX := 0
	maxY := 0
	for x := 1; x <= board.xMax-size; x++ {
		for y := 1; y <= board.yMax-size; y++ {
			pw := squarePower(board, x, y, size, size)
			if pw > maxPower {
				maxPower = pw
				maxX = x
				maxY = y
			}
		}
	}
	return maxX, maxY, maxPower
}

func squarePower(board Board, x int, y int, xSize int, ySize int) int {
	power := 0
	toX := x + xSize
	if toX > board.xMax {
		toX = board.xMax
	}
	if ySize > board.yMax {
		ySize = board.yMax
	}
	toY := y + ySize
	for i := x; i < toX; i++ {
		for j := y; j < toY; j++ {
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
	x, y, power := findMaxSquare(board, 3)
	println("part1:", x, y, power)
	maxPower, maxY, maxX, maxSize := findVariableSizeMaxSquare(board)
	println("part2:", maxX, maxY, maxSize, maxPower)
}

func findVariableSizeMaxSquare(board Board) (int, int, int, int) {
	maxPower := 0
	maxY := 0
	maxX := 0
	maxSize := 0
	for size := 1; size <= 300; size++ {
		x, y, power := findMaxSquare(board, size)
		if power > maxPower {
			maxPower = power
			maxSize = size
			maxX = x
			maxY = y
		}
	}
	return maxPower, maxY, maxX, maxSize
}
