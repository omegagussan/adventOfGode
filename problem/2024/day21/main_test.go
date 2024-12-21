package main

import (
	"testing"
)

func TestGoToVertical(t *testing.T) {
	res := goToNumerical('3', '9')
	if res != "vv" {
		t.Errorf("Expected vv, got %s", res)
	}
}

func TestGoToHorizontal(t *testing.T) {
	res := goToNumerical('3', '1')
	if res != ">>" {
		t.Errorf("Expected >>, got %s", res)
	}
}

func TestGoToDiagonal(t *testing.T) {
	res := goToNumerical('1', '9')
	if res != "vv<<" {
		t.Errorf("Expected vv<<, got %s", res)
	}
}

func TestGoToDiagonal2(t *testing.T) {
	res := goToNumerical('9', '1')
	if res != ">>^^" {
		t.Errorf("Expected vv>>, got %s", res)
	}
}

func TestGoToA(t *testing.T) {
	res := goToNumerical('A', '1')
	if res != ">>v" {
		t.Errorf("Expected >>v, got %s", res)
	}
}

func TestGoTo0(t *testing.T) {
	res := goToNumerical('0', '2')
	if res != "v" {
		t.Errorf("Expected v, got %s", res)
	}
}

func TestGoFrom0(t *testing.T) {
	res := goToNumerical('2', '0')
	if res != "^" {
		t.Errorf("Expected ^, got %s", res)
	}
}

func TestGoToA2(t *testing.T) {
	res := goToNumerical('A', '9')
	if res != "vvv" {
		t.Errorf("Expected vvv, got %s", res)
	}
}

func TestGoToFromA(t *testing.T) {
	res := goToNumerical('1', 'A')
	if res != "^<<" {
		t.Errorf("Expected ^<<, got %s", res)
	}
}

func TestGoToFrom2(t *testing.T) {
	res := goToNumerical('0', 'A')
	if res != "<" {
		t.Errorf("Expected <, got %s", res)
	}
}

func TestNumericalToDirectional(t *testing.T) {
	res := numericalToDirectional("029A")
	if res != "<A^A>^^AvvvA" && res != "<A^A^>^AvvvA" && res != "<A^A^^>AvvvA" {
		t.Errorf("Expected something else, got %s", res)
	}
}

func TestDirectionalToDirectional(t *testing.T) {
	res := directionalToDirectional("<A^A^^>AvvvA")
	if len(res) != 28 {
		t.Errorf("Expected %d got %d", len(res), 28)
	}
}

func TestDirectionalToDirectional2(t *testing.T) {
	res := directionalToDirectional("v<<A>>^A<A>AvA<^AA>A<vAAA>^A")
	if len(res) != 68 {
		t.Errorf("Expected %d got %d", len(res), 68)
	}
}
