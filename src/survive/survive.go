package main

import(
	"fmt"
	//"engine"
	"gameio"
)

var(
	continueRunning = true
)

func main(){
	gameio.ClearConsole()
	gameio.PrintGameStart()

	for continueRunning {
		fmt.Println("Continue running top of loop")
		var option string
		fmt.Scanln(&option)

		gameio.HandleGameInput(option, &continueRunning)

		fmt.Println("Coninue running...", continueRunning)
	}

	gameio.PrintThanksForPlaying()
}