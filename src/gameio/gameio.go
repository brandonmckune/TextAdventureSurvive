package gameio

import(
	"fmt"
	"strings"
	"engine"
)

var(
	actions = make(map[uint64]func(*bool))
)

func init(){
	buildActionsMap()
}

func buildActionsMap(){
	actions[engine.GetHash("n")] = newGame
	actions[engine.GetHash("q")] = quitGame 
}

func newGame(_ *bool){
	
}

func quitGame(continueRunning *bool){
	*continueRunning = false
}

func HandleGameInput(input string, continueRunning *bool){
	input = strings.ToLower(input)

	if action, found := actions[engine.GetHash(input)]; found{
		action(continueRunning)
	}else{
		//TODO: add logging for incorrect functionality
		fmt.Println("No action found....")
	}
}