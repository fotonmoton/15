package main

import "os"

func main() {

	message := "Only 'game' and 'solve' options supported"

	if len(os.Args) == 1 {
		panic(message)
	}

	switch os.Args[1] {
	case "game":
		StartGame()
	case "solve":
		StartSolver()
	default:
		panic(message)
	}
}
