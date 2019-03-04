package engine

import (
	"fmt"
	"strings"
	//"strconv"
)

const(
	ACTION_DELIMITER = " "
	NO_ACTION = ""
)

var (
	generalActions = make(map[uint64]func(*GameDetails, *bool, ...string))
)

func init() {
	buildActionsMap()
}

func buildActionsMap() {
	generalActions[GetHash("n")] = newGame
	generalActions[GetHash("l")] = loadGame
	generalActions[GetHash("h")] = displayHelp
	generalActions[GetHash("q")] = quitGame
	generalActions[GetHash("north")] = goNorth
	generalActions[GetHash("south")] = goSouth
	generalActions[GetHash("east")] = goEast
	generalActions[GetHash("west")] = goWest
}

func newGame(game *GameDetails, _ *bool, _ ... string) {
	game.NewGame()
}

func loadGame(_ *GameDetails, _ *bool, _ ... string) {
	fmt.Println("Load Game is not currently available.")
}

func displayHelp(game *GameDetails, _ *bool, _ ... string) {
	game.DisplayScreen("Help")
}

func quitGame(_ *GameDetails, continueRunning *bool, _ ... string) {
	*continueRunning = false
}

func goNorth(game *GameDetails, _ *bool, vars ... string){
	game.MoveNorth(1)
}

func goSouth(game *GameDetails, _ *bool, vars ... string){
	game.MoveSouth(1)
}

func goEast(game *GameDetails, _ *bool, vars ... string){
	game.MoveEast(1)
}

func goWest(game *GameDetails, _ *bool, vars ... string){
	game.MoveWest(1)
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

	parts := strings.Split(input, ACTION_DELIMITER)
		
	switch len(parts) {
	case 1:
		//general action defined by function
		if action, found := generalActions[GetHash(input)]; found {
			action(game, continueRunning, NO_ACTION)
		} else {
			//TODO: add logging for incorrect functionality
			fmt.Println("Command length 1: No action found....[", input, "]")
		}
	case 2: //move functions
		if action, found := generalActions[GetHash(parts[0])]; found {
			action(game, continueRunning, parts[1])
		} else {
			//TODO: add logging for incorrect functionality
			fmt.Println("Command length 2: No action found....")
		}
	default:
		fmt.Println("Cannot take action: ", input)
	}
}
