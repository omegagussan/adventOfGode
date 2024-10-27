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
			pw := squarePower(board, x, y, size)
			if pw > maxPower {
				maxPower = pw
				maxX = x
				maxY = y
			}
		}
	}
	return maxX, maxY, maxPower
}

func findMaxSquareCached(board Board, size int, cache PowerCache) (int, int, int) {
	maxPower := 0
	maxX := 0
	maxY := 0
	for x := 1; x <= board.xMax-size; x++ {
		for y := 1; y <= board.yMax-size; y++ {
			pw := squarePowerCached(board, x, y, size, cache)
			if pw > maxPower {
				maxPower = pw
				maxX = x
				maxY = y
			}
		}
	}
	return maxX, maxY, maxPower
}

func squarePower(board Board, x int, y int, size int) int {
	power := 0
	toX := x + size
	if toX > board.xMax {
		toX = board.xMax
	}
	toY := y + size
	if toY > board.yMax {
		toY = board.yMax
	}
	for i := x; i < toX; i++ {
		for j := y; j < toY; j++ {
			power += board.board[i][j]
		}
	}
	return power
}

func squarePowerCached(board Board, x int, y int, size int, cache PowerCache) int {
	if size == 1 {
		va := board.board[x][y]
		cache[PowerCacheKey{x, y, 1}] = PowerKeyValue{value: va, exists: true}
		return va
	}
	power := cache[PowerCacheKey{x, y, size - 1}]

	if power.exists == false {
		panic("this should not happen")
	}

	powerVal := power.value

	for i := x; i < x+size-1; i++ {
		powerVal += board.board[i][(y + size - 1)]
	}
	for i := y; i < y+size-1; i++ {
		powerVal += board.board[(x + size - 1)][i]
	}
	powerVal += board.board[(x + size - 1)][(y + size - 1)]

	cache[PowerCacheKey{x, y, size}] = PowerKeyValue{value: powerVal, exists: true}
	return powerVal
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

type PowerCacheKey struct {
	x    int
	y    int
	size int
}

type PowerKeyValue struct {
	exists bool
	value  int
}

type PowerCache map[PowerCacheKey]PowerKeyValue

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

	//PowerCacheKey:power cache
	cache := make(map[PowerCacheKey]PowerKeyValue)
	//obs! This is an assumption
	for size := 1; size <= 30; size++ {
		x, y, power := findMaxSquareCached(board, size, cache)
		if power > maxPower {
			maxPower = power
			maxSize = size
			maxX = x
			maxY = y
		}
	}
	return maxPower, maxY, maxX, maxSize
}
