package lib

import (
	"fmt"
	"math"
	"math/rand"

	"golang.org/x/exp/slices"
)

type Board struct {
	grid  [16]int
	empty [2]int
}

const ROW_COUNT = 4

var SOLVED_GRID = [16]int{
	1, 2, 3, 4,
	5, 6, 7, 8,
	9, 10, 11, 12,
	13, 14, 15, 0,
}

type Direction int

const (
	LEFT Direction = iota
	RIGHT
	UP
	DOWN
)

func NewBoard() *Board {
	return &Board{
		grid:  [16]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 0},
		empty: [2]int{3, 3},
	}
}

// If all pieces on their desired places
// no more moves are needed and we can say that
// board is solved.
func (board *Board) Solved() bool {
	return board.neededMoves() == 0
}

// Faster way to check if board is solved.
// Arrays are comparable in Go so we can simply
// compare desired state with current
func (board *Board) SolvedFast() bool {
	return board.grid == SOLVED_GRID
}

func (b *Board) Print() {
	for i, cell := range b.grid {
		if cell == 0 {
			fmt.Printf("   ")
		} else {
			fmt.Printf("%3d", cell)
		}
		if i == 3 || i == 7 || i == 11 {
			fmt.Println()
		}
	}
	fmt.Println()
}

func (b *Board) PossibleDirections() []Direction {
	directions := []Direction{}

	if b.empty[0] != 0 {
		directions = append(directions, UP)
	}

	if b.empty[0] != 3 {
		directions = append(directions, DOWN)
	}

	if b.empty[1] != 0 {
		directions = append(directions, LEFT)
	}

	if b.empty[1] != 3 {
		directions = append(directions, RIGHT)
	}

	return directions
}

// "Moves" empty cell to new position.
// It's easier to reason if we will move
// empty cell as another piece rather than
// moving pieces that are surrounds it.
func (b *Board) Move(d Direction) {
	possibleDirections := b.PossibleDirections()

	if slices.Index(possibleDirections, d) == -1 {
		return
	}

	toRow, toCol := directionToStep(d)

	newRow := b.empty[0] + toRow
	newCol := b.empty[1] + toCol

	piceToSwap := b.get(newRow, newCol)

	b.set(b.empty[0], b.empty[1], piceToSwap)
	b.set(newRow, newCol, 0)

	b.empty[0] = newRow
	b.empty[1] = newCol
}

func (b *Board) Shuffle(steps int) []Direction {
	moves := []Direction{}

	for i := 0; i < steps; i++ {
		possibleDirections := b.PossibleDirections()

		// Remove opposite moves to prevent
		// moving around one cell
		if len(moves) != 0 {
			last := moves[len(moves)-1]
			possibleDirections = slices.DeleteFunc(
				possibleDirections,
				func(direction Direction) bool {
					return oppositeDirections(direction, last)
				})
		}

		rand.Shuffle(len(possibleDirections), func(i, j int) {
			possibleDirections[i], possibleDirections[j] = possibleDirections[j], possibleDirections[i]
		})

		nextMove := possibleDirections[0]

		b.Move(nextMove)

		moves = append(moves, nextMove)
	}

	return moves
}

// Optimistic number. Indicates
// sum of number of moves each board piece should do
// to get to desired position. It ignores real "circular"
// moves and calculates moves as if only one piece exists on the board.
func (board *Board) neededMoves() int {
	neededMoves := 0

	for row := 0; row < 4; row++ {
		for col := 0; col < 4; col++ {

			number := board.get(row, col)

			if number == 0 {
				continue
			}

			neededMoves += rectilinearDistance(number, row, col)
		}
	}

	return neededMoves
}

func (b *Board) get(row, col int) int {
	return b.grid[row*ROW_COUNT+col]
}

func (b *Board) set(row, col, val int) {
	b.grid[row*ROW_COUNT+col] = val
}

func originalPosition(number int) (int, int) {
	return (number - 1) / 4, (number - 1) % 4
}

// Or "Manhattan distance". We use it to calculate "shortest" path
// to desired piece position.
// https://en.wikipedia.org/wiki/Taxicab_geometry
func rectilinearDistance(number, i, j int) int {
	origRow, origCol := originalPosition(number)
	return int(math.Abs(float64(origRow-i)) + math.Abs(float64(origCol-j)))
}

func directionToStep(d Direction) (int, int) {
	switch d {
	case UP:
		return -1, 0
	case DOWN:
		return 1, 0
	case LEFT:
		return 0, -1
	case RIGHT:
		return 0, 1
	default:
		return 0, 0
	}
}

func oppositeDirections(a Direction, b Direction) bool {
	ar, al := directionToStep(a)
	br, bl := directionToStep(b)

	return ar+br == 0 && al+bl == 0
}
