package main

import (
	"engine"
)

const (
	STARTSCREEN = "0"
)

var (
	continueRunning = true
	action          string
	game            engine.GameDetails
)

func init() {
	engine.ClearConsole()
	engine.LoadDefaultData(&game)
	engine.ClearConsole()
}

func main() {
	game.DisplayScreen(STARTSCREEN)
	engine.RetreiveAndHandleGameInput(&game, &continueRunning)

	for continueRunning {

		engine.RetreiveAndHandleGameInput(&game, &continueRunning)
	}
}
