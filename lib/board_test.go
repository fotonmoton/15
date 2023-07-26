package lib

import (
	"testing"

	"golang.org/x/exp/slices"
)

func TestSolvedState(t *testing.T) {
	board := NewBoard()
	if !board.Solved() {
		t.Error("Initial state should be solved")
	}

	boardWithShuffledPieces := Board{grid: [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 0, 15}, empty: [2]int{3, 3}}

	if boardWithShuffledPieces.Solved() {
		t.Error("Shuffled board should not be solved")
	}
}

func TestSolvedFast(t *testing.T) {
	board := NewBoard()
	if !board.SolvedFast() {
		t.Error("Initial state should be solved")
	}

	boardWithShuffledPieces := Board{
		grid: [16]int{
			1, 2, 3, 4,
			5, 6, 7, 8,
			9, 10, 11, 12,
			13, 14, 0, 15,
		},
		empty: [2]int{3, 3}}

	if boardWithShuffledPieces.SolvedFast() {
		t.Error("Shuffled board should not be solved")
	}
}

func TestPossibleDirections(t *testing.T) {
	board := NewBoard()
	directions := board.PossibleDirections()

	if len(directions) != 2 {
		t.Error("For initial state only UP and LEFT directions should be available")
	}

	isUPresent := slices.Index(directions, UP)
	isLeftPresent := slices.Index(directions, LEFT)

	if isLeftPresent == -1 || isUPresent == -1 {
		t.Error("UP and LEFT directions should be present")
	}

	board.Move(LEFT)

	directions = board.PossibleDirections()

	if len(directions) != 3 {
		t.Error("should be 3 possible directions after one move left from initial state")
	}
}

func TestMove(t *testing.T) {
	board := NewBoard()
	board.Move(LEFT)

	toTheRight := board.grid == [16]int{
		1, 2, 3, 4,
		5, 6, 7, 8,
		9, 10, 11, 12,
		13, 14, 0, 15,
	}

	if !toTheRight {
		t.Error("Move should move pieces")
	}

	if board.empty != [2]int{3, 2} {
		t.Error("after Move new empty position should be set")
	}

	board.Move(UP)
	board.Move(UP)
	board.Move(UP)

	tripleUp := board.grid == [16]int{
		1, 2, 0, 4,
		5, 6, 3, 8,
		9, 10, 7, 12,
		13, 14, 11, 15,
	}

	if !tripleUp {
		t.Error("Corrupt state after moving")
	}

	board.Move(UP)

	asBefore := board.grid == [16]int{
		1, 2, 0, 4,
		5, 6, 3, 8,
		9, 10, 7, 12,
		13, 14, 11, 15,
	}

	if !asBefore {
		t.Error("If we cannot move further state should stay the same")
	}
}

func TestOppositeDirections(t *testing.T) {
	vertical := oppositeDirections(UP, DOWN)
	horizontal := oppositeDirections(LEFT, RIGHT)

	if !vertical || !horizontal {
		t.Error("Opposite direction should return true")
	}
}

func TestShuffle(t *testing.T) {
	board := NewBoard()
	board.Shuffle(2)

	if board.SolvedFast() {
		t.Error("Board should be in unsolved state after shuffle")
	}
}
