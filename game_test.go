package main

import (
	"testing"
)

func TestNewGame(t *testing.T) {
	game := New(1, 2)

	if len(game.Grid) != 1 {
		t.Fail()
	}

	if len(game.Grid[0]) != 2 {
		t.Fail()
	}
}

func TestValorPadrao(t *testing.T) {
	game := New(1, 1)

	if game.Grid[0][0] {
		t.Fail()
	}
}

func TestSeed(t *testing.T) {
	game := New(2, 2)

	game.Seed(
		[][]bool{
			[]bool{true, false},
			[]bool{false, true},
		})

	if !game.Grid[0][0] {
		t.Error("[0][0]")
	}
	if game.Grid[0][1] {
		t.Error("[0][1]")
	}
	if game.Grid[1][0] {
		t.Error("[1][0]")
	}
	if !game.Grid[1][1] {
		t.Error("[1][1]")
	}
}

func TestUnderPopulation(t *testing.T) {
	game := New(2, 2)
	game.Seed(
		[][]bool{
			[]bool{true, false},
			[]bool{false, false},
		})

	game.Step()

	if game.Grid[0][0] {
		t.Fail()
	}
}

func TestOverCrowding(t *testing.T) {
	game := New(2, 3)
	game.Seed(
		[][]bool{
			[]bool{true, true, true},
			[]bool{true, true, true},
		})

	game.Step()

	if !game.Grid[0][0] {
		t.Error("[0][0]")
	}
	if game.Grid[0][1] {
		t.Error("[0][1]")
	}
	if !game.Grid[0][2] {
		t.Error("[0][2]")
	}
	if !game.Grid[1][0] {
		t.Error("[1][0]")
	}
	if game.Grid[1][1] {
		t.Error("[1][1]")
	}
	if !game.Grid[1][2] {
		t.Error("[1][2]")
	}
}

func TestReproduction(t *testing.T) {
	game := New(2, 3)
	game.Seed(
		[][]bool{
			[]bool{true, false, true},
			[]bool{false, true, false},
		})

	game.Step()

	if !game.Grid[0][1] {
		t.Error("[0][1]")
	}
}

func TestVivo2Vizinhos(t *testing.T) {
	game := New(2, 3)
	game.Seed(
		[][]bool{
			[]bool{true, true, true},
			[]bool{false, false, false},
		})

	game.Step()

	if !game.Grid[0][1] {
		t.Error("[0][1]")
	}
}

func TestMorto2Vizinhos(t *testing.T) {
	game := New(2, 3)
	game.Seed(
		[][]bool{
			[]bool{true, false, true},
			[]bool{false, false, false},
		})

	game.Step()

	if game.Grid[0][1] {
		t.Error("[0][1]")
	}
}
