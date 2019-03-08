package engine

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	//"strconv"
)

const (
	ACTION_DELIMITER = " "
	NO_ACTION        = ""
	ACTION_MOVE      = "MOVE"
	ACTION_RUN       = "RUN"
	ACTION_LOOK      = "LOOK"
	ACTION_HIDE      = "HIDE"
)

var (
	generalActions = make(map[uint64]func(*GameDetails, *bool, []string))
)

func init() {
	buildActionsMap()
}

func buildActionsMap() {
	generalActions[GetHash("N")] = newGame
	generalActions[GetHash("L")] = loadGame
	generalActions[GetHash("H")] = displayHelp
	generalActions[GetHash("Q")] = quitGame
	generalActions[GetHash(ACTION_MOVE)] = move
	generalActions[GetHash(ACTION_RUN)] = run
}

func newGame(game *GameDetails, _ *bool, _ []string) {
	game.NewGame()
}

func loadGame(_ *GameDetails, _ *bool, _ []string) {
	fmt.Println("Load Game is not currently available.")
}

func displayHelp(game *GameDetails, _ *bool, _ []string) {
	game.DisplayScreen("Help")
}

func quitGame(_ *GameDetails, continueRunning *bool, _ []string) {
	*continueRunning = false
}

func run(game *GameDetails, _ *bool, vars []string) {
	switch len(vars) {
	case 1:
		game.Move(vars[0], 2)
	case 2:
		tileNumber, err := strconv.Atoi(vars[1])
		check(err)
		game.Move(vars[0], tileNumber)
	default:
		fmt.Println("Invalid command. [run ", vars, "]")
	}
}

func move(game *GameDetails, _ *bool, vars []string) {
	switch len(vars) {
	case 1:
		game.Move(vars[0], 1)
	case 2:
		tileNumber, err := strconv.Atoi(vars[1])
		check(err)
		game.Move(vars[0], tileNumber)
	default:
		fmt.Println("Invalid command. [move ", vars, "]")
	}
}

func RetreiveAndHandleGameInput(game *GameDetails, continueRunning *bool) {
	var actions string

	actions = RetreiveGameInput()
	HandleGameInput(actions, game, continueRunning)
}

func RetreiveGameInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	var userInput string

	scanner.Scan()
	userInput += scanner.Text()

	if err := scanner.Err(); err != nil {
		os.Exit(1)
	}

	return userInput
}

func HandleGameInput(input string, game *GameDetails, continueRunning *bool) {
	input = strings.ToUpper(strings.TrimSpace(input))
	parts := strings.Split(input, ACTION_DELIMITER)

	switch len(parts) {
	case 1:
		//general action defined by function
		if action, found := generalActions[GetHash(parts[0])]; found {
			action(game, continueRunning, nil)
		} else {
			//TODO: add logging for incorrect functionality
			fmt.Println("Command length 1: No action found....[", parts, "]")
		}
	default:
		if action, found := generalActions[GetHash(parts[0])]; found {
			action(game, continueRunning, parts[1:])
		} else {
			//TODO: add logging for incorrect functionality
			fmt.Println("Action[", input, "] not found....")
		}
	}
}
