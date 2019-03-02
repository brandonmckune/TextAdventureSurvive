package main

import (
	"fmt"

	//"engine"

	"bufio"
	"gameio"
	"io"
	"os"
)

var (
	continueRunning = true
	option          string
	buffer          = make([]byte, 40)
)

// func main(){

// 	fmt.Println("Opening file")

// 	file, err := os.Open("../data.wad")
// 	check(err)
// 	reader := bufio.NewReader(file)
// 	_, err = io.ReadFull(reader, buffer)
// 	check(err)

// 	for err != io.EOF {
// 		fmt.Println(buffer)
// 		_, err = io.ReadFull(reader, buffer)
// 	}

// 	fmt.Scanln(&option)

// }

func check(e error) {
	if e != nil {
		fmt.Println(e)
		fmt.Scanln(&option)
		panic(e)
	}
}

func main() {
	gameio.ClearConsole()

	fmt.Println("Opening file")
	fmt.Println("file path: ", os.Args[0])
	var filepath = os.Getenv("GOPATH") + "/bin/data.wad"
	file, err := os.OpenFile(filepath, os.O_RDONLY, 0755)
	check(err)
	reader := bufio.NewReader(file)
	//_, err = io.ReadFull(reader, buffer)
	myLine, err := reader.ReadString('\n')
	check(err)

	for err != io.EOF {
		fmt.Print(myLine)
		myLine, err = reader.ReadString('\n')
		//fmt.Println(string(buffer[:]))
		//fmt.Printf("%c", buffer)
		//_, err = io.ReadFull(reader, buffer)
	}

	fmt.Scanln(&option)

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
