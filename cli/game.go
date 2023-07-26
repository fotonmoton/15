package main

import (
	"15/lib"
	"fmt"
	"time"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

type CliGame struct {
	board *lib.Board
}

func (g *CliGame) PrintState() {
	// Works only on Linux
	fmt.Print("\033[H\033[2J")
	fmt.Printf("To quit game press ESC.\n")
	g.board.Print()
	fmt.Printf("Solved: %t\n", g.board.SolvedFast())
}

func (g *CliGame) Loop() {
	g.PrintState()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			return true, nil
		case keys.Up:
			g.board.Move(lib.DOWN)
		case keys.Down:
			g.board.Move(lib.UP)
		case keys.Left:
			g.board.Move(lib.RIGHT)
		case keys.Right:
			g.board.Move(lib.LEFT)
		default:
			fmt.Printf("\rYou pressed: %s\n", key)
		}

		g.PrintState()

		if g.board.SolvedFast() {
			fmt.Printf("\rYou Won!\n")
			return true, nil
		}

		return false, nil
	})
}

func (g *CliGame) Solve() {

	path, cost := lib.Solver(g.board)

	reverse := []*lib.Board{}
	path.ForEach(func(b *lib.Board) {
		reverse = append([]*lib.Board{b}, reverse...)
	})

	for _, b := range reverse {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Cost is: %d\n", cost)
		b.Print()
		time.Sleep(1 * time.Second)
	}
}

func StartGame() {
	game := CliGame{board: lib.NewBoard()}

	game.board.Shuffle(10)

	game.Loop()
}

func StartSolver() {
	game := CliGame{board: lib.NewBoard()}

	game.board.Shuffle(20)

	game.Solve()
}
