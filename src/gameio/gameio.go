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
	fmt.Println("You've created a new game")
}

func quitGame(continueRunning *bool){
	fmt.Println("You are a quiter? Okay...")
	*continueRunning = false
	
	fmt.Println(*continueRunning)
}

func HandleGameInput(input string, continueRunning *bool){
	input = strings.ToLower(input)

	fmt.Println("Inside the HandleGameInput: ", input, " | ", *continueRunning)

	if action, found := actions[engine.GetHash(input)]; found{
		action(continueRunning)
	}else{
		fmt.Println("No action found....")
	}

	fmt.Println("Post handle input: ", *continueRunning)
}