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
	fakeInput       string
)

func init() {
	engine.ClearConsole()
	engine.LoadDefaultData(&game)
	engine.ClearConsole()
	game.DisplayScreen(STARTSCREEN)
}

func main() {
	engine.HandleGameInput("n", &game, &continueRunning)

	for continueRunning {
		engine.RetreiveAndHandleGameInput(&game, &continueRunning)
	}
}
