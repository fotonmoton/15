package lib

import "testing"

func TestSolved(t *testing.T) {
	board := NewBoard()

	path, cost := Solver(board)

	if cost != 0 {
		t.Error("cost for solved board should be 0")
	}

	if *path.Peek() != *board {
		t.Error("root board should be in path")
	}

	if path.Size() != 1 {
		t.Error("only one board should be in path")
	}
}

func TestSimpleShuffle(t *testing.T) {
	board := NewBoard()

	board.Shuffle(10)

	path, cost := Solver(board)

	if cost > 10 {
		t.Error("it should be more greedy to find optimal path")
	}

	if path.Pop().grid != SOLVED_GRID {
		t.Error("last board should be solved")
	}
}
