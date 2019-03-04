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
	fakeInput		string
	isDebug			= false
)

func init() {
	engine.ClearConsole()
	engine.LoadDefaultData(&game)
	engine.ClearConsole()
	game.DisplayScreen(STARTSCREEN)
	fakeInput = "n"
}

func main() {

	if isDebug {
		engine.RetreiveAndHandleGameInput(&game, &continueRunning, &fakeInput)
	}

	for continueRunning {
		engine.RetreiveAndHandleGameInput(&game, &continueRunning, nil)
	}
}
