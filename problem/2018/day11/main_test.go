package main

import "testing"

func TestPowerLevelWithPositiveSerial(t *testing.T) {
	result := powerLevel(57, 122, 79)
	expected := -5
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPowerLevelWithPositiveSerial2(t *testing.T) {
	result := powerLevel(39, 217, 196)
	expected := 0
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestPowerLevelWithPositiveSerial3(t *testing.T) {
	result := powerLevel(8, 3, 5)
	expected := 4
	if result != expected {
		t.Errorf("Expected %d, but got %d", expected, result)
	}
}

func TestMakeBoardCreatesCorrectSize(t *testing.T) {
	board := makeBoard(18, 300, 300)
	if len(board.board) != 300 {
		t.Errorf("Expected board size %d, but got %d", 300, len(board.board))
	}
	if len(board.board[1]) != 300 {
		t.Errorf("Expected board size %d, but got %d", 300, len(board.board[1]))

	}
	if board.xMax != 300 {
		t.Errorf("Expected xMax %d, but got %d", 300, board.xMax)
	}
	if board.yMax != 300 {
		t.Errorf("Expected yMax %d, but got %d", 300, board.yMax)
	}
}

func TestFindMaxSquareReturnsCorrectValues(t *testing.T) {
	board := makeBoard(18, 300, 300)
	x, y, power := findMaxSquare(board)
	if x != 33 || y != 45 || power != 29 {
		t.Errorf("Expected (33, 45, 29), but got (%d, %d, %d)", x, y, power)
	}
}

func TestFindMaxSquareReturnsCorrectValues2(t *testing.T) {
	board := makeBoard(42, 300, 300)
	x, y, power := findMaxSquare(board)
	if x != 21 || y != 61 || power != 30 {
		t.Errorf("Expected (21, 61, 30), but got (%d, %d, %d)", x, y, power)
	}
}
