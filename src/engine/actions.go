package engine

import (
	"fmt"
	"strings"
)

var (
	actions = make(map[uint64]func(*bool))
)

func init() {
	buildActionsMap()
}

func buildActionsMap() {
	actions[GetHash("n")] = newGame
	actions[GetHash("q")] = quitGame
}

func newGame(_ *bool) {

}

func quitGame(continueRunning *bool) {
	*continueRunning = false
}

func RetreiveAndHandleGameInput(game *GameDetails, continueRunning *bool) {
	action := RetreiveGameInput()
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
		action(continueRunning)
	} else {
		//TODO: add logging for incorrect functionality
		fmt.Println("No action found....")
	}
}
