package main

type Board struct {
	xMax, yMax int
	board      map[int]map[int]int
}

type PowerCacheKey struct {
	x, y, size int
}

type PowerKeyValue struct {
	exists bool
	value  int
}

type PowerCache map[PowerCacheKey]PowerKeyValue

func makeBoard(serialNr, xMax, yMax int) Board {
	board := make(map[int]map[int]int)
	for x := 1; x <= xMax; x++ {
		board[x] = make(map[int]int)
		for y := 1; y <= yMax; y++ {
			board[x][y] = powerLevel(serialNr, x, y)
		}
	}
	return Board{xMax, yMax, board}
}

func findMaxSquare(board Board, size int) (int, int, int) {
	maxPower, maxX, maxY := 0, 0, 0
	for x := 1; x <= board.xMax-size; x++ {
		for y := 1; y <= board.yMax-size; y++ {
			pw := squarePower(board, x, y, size)
			if pw > maxPower {
				maxPower, maxX, maxY = pw, x, y
			}
		}
	}
	return maxX, maxY, maxPower
}

func findMaxSquareCached(board Board, size int, cache PowerCache) (int, int, int) {
	maxPower, maxX, maxY := 0, 0, 0
	for x := 1; x <= board.xMax-size; x++ {
		for y := 1; y <= board.yMax-size; y++ {
			pw := squarePowerCached(board, x, y, size, cache)
			if pw > maxPower {
				maxPower, maxX, maxY = pw, x, y
			}
		}
	}
	return maxX, maxY, maxPower
}

func squarePower(board Board, x, y, size int) int {
	power := 0
	toX, toY := min(x+size, board.xMax), min(y+size, board.yMax)
	for i := x; i < toX; i++ {
		for j := y; j < toY; j++ {
			power += board.board[i][j]
		}
	}
	return power
}

func squarePowerCached(board Board, x, y, size int, cache PowerCache) int {
	if size == 1 {
		val := board.board[x][y]
		cache[PowerCacheKey{x, y, 1}] = PowerKeyValue{true, val}
		return val
	}
	power := cache[PowerCacheKey{x, y, size - 1}]
	if !power.exists {
		panic("this should not happen")
	}
	powerVal := power.value
	for i := x; i < x+size-1; i++ {
		powerVal += board.board[i][y+size-1]
	}
	for i := y; i < y+size-1; i++ {
		powerVal += board.board[x+size-1][i]
	}
	powerVal += board.board[x+size-1][y+size-1]
	cache[PowerCacheKey{x, y, size}] = PowerKeyValue{true, powerVal}
	return powerVal
}

func powerLevel(serialNr, x, y int) int {
	rackId := x + 10
	power := ((rackId*y + serialNr) * rackId / 100 % 10) - 5
	return power
}

func findVariableSizeMaxSquare(board Board) (int, int, int, int) {
	maxPower, maxX, maxY, maxSize := 0, 0, 0, 0
	cache := make(PowerCache)
	for size := 1; size <= 30; size++ {
		x, y, power := findMaxSquareCached(board, size, cache)
		if power > maxPower {
			maxPower, maxX, maxY, maxSize = power, x, y, size
		}
	}
	return maxPower, maxX, maxY, maxSize
}

func main() {
	board := makeBoard(9306, 300, 300)
	x, y, power := findMaxSquare(board, 3)
	println("part1:", x, y, power)
	maxPower, maxX, maxY, maxSize := findVariableSizeMaxSquare(board)
	println("part2:", maxX, maxY, maxSize, maxPower)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
