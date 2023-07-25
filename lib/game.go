package lib

import (
	"fmt"

	"atomicgo.dev/keyboard"
	"atomicgo.dev/keyboard/keys"
)

type Game struct {
	board *Board
}

func (g *Game) PrintState() {
	// Works only on Linux
	fmt.Print("\033[H\033[2J")
	fmt.Printf("To quit game press ESC.\n")
	g.board.Print()
	fmt.Printf("Solved: %t\n", g.board.SolvedFast())
}

func (g *Game) Loop() {
	g.PrintState()

	keyboard.Listen(func(key keys.Key) (stop bool, err error) {
		switch key.Code {
		case keys.CtrlC, keys.Escape:
			return true, nil
		case keys.Up:
			g.board.Move(DOWN)
		case keys.Down:
			g.board.Move(UP)
		case keys.Left:
			g.board.Move(RIGHT)
		case keys.Right:
			g.board.Move(LEFT)
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

func Start() {
	game := Game{board: NewBoard()}

	game.board.Shuffle(10)

	game.Loop()
}
