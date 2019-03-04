package engine

import (
	"fmt"
	"strings"
)

var (
	actions = make(map[uint64]func(*GameDetails, *bool))
)

func init() {
	buildActionsMap()
}

func buildActionsMap() {
	actions[GetHash("n")] = newGame
	actions[GetHash("l")] = loadGame
	actions[GetHash("h")] = displayHelp
	actions[GetHash("q")] = quitGame
	actions[GetHash("north")] = goNorth
}

func newGame(game *GameDetails, _ *bool) {
	game.NewGame()
}

func loadGame(_ *GameDetails, _ *bool) {
	fmt.Println("Load Game is not currently available.")
}

func displayHelp(game *GameDetails, _ *bool) {
	game.DisplayScreen("Help")
}

func quitGame(_ *GameDetails, continueRunning *bool) {
	*continueRunning = false
}

func goNorth(game *GameDetails, _ *bool){
	game.MoveNorth(1)
}

func RetreiveAndHandleGameInput(game *GameDetails, continueRunning *bool, option *string) {
	var action string

	if option == nil {
		action = RetreiveGameInput()
	} else{
		action = *option
	}
	HandleGameInput(action, game, continueRunning)
}

func RetreiveGameInput() string {
	var input string
	fmt.Scanln(&input)
	return input
}

func HandleGameInput(input string, game *GameDetails, continueRunning *bool) {
	input = strings.ToLower(input)

	if action, found := actions[GetHash(input)]; found {
		action(game, continueRunning)
	} else {
		//TODO: add logging for incorrect functionality
		fmt.Println("No action found....")
	}
}
