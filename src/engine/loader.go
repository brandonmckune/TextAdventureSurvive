package engine

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

const (
	LEVELS    = "/bin/levels.wad"
	SCREENS   = "/bin/screens.wad"
	DELIMITER = "::"
	SCREEN    = "screen"
	LEVEL     = "level"
)

var (
	screensFilePath string
	levelsFilePath  string
)

func check(e error) {
	if e != nil && e != io.EOF {
		fmt.Println(e)
		panic(e)
	}
}

func init() {
	levelsFilePath = os.Getenv("GOPATH") + LEVELS
	screensFilePath = os.Getenv("GOPATH") + SCREENS
}

func LoadDefaultData(game *GameDetails) {
	if game == nil {
		game = new(GameDetails)
	}

	fmt.Println("Game loading ")
	loadScreens(game)
	loadLevels(game)
	fmt.Println("Game finished loading.")
}

func loadScreens(game *GameDetails) {
	fmt.Print("Loading screens ")
	file, err := os.OpenFile(screensFilePath, os.O_RDONLY, 0755)
	check(err)
	reader := bufio.NewReader(file)

	myLine, err := reader.ReadString('\n')
	check(err)

	/* Steps for Loading wad into game
	 * 1. Read a line
	 * 2. Determine its location (Screen vs Level)
	 * 3. Update Object
	 */
	for err != io.EOF {
		if strings.Contains(myLine, DELIMITER) {
			fmt.Print(".") // status load indicator
			parts := strings.Split(myLine, DELIMITER)

			if len(parts) > 3 {
				panic("Failed to load screens.wad file.")
			}

			screenSequence := parts[0] //sequence key first item defined
			newScreen := new(ScreenDetails)

			myLine, err = reader.ReadString('\n')
			check(err)

			//Buid out the screen
			for idx := 0; err != io.EOF && !strings.Contains(myLine, DELIMITER); idx++ {
				newScreen.LoadRow(myLine, idx)
				myLine, err = reader.ReadString('\n')
				check(err)
			}

			game.AddScreen(screenSequence, newScreen)
		}
	}

	fmt.Println("finished successfully")
}

func loadLevels(game *GameDetails) {
	fmt.Print("Loading levels ")
	file, err := os.OpenFile(levelsFilePath, os.O_RDONLY, 0755)
	check(err)
	reader := bufio.NewReader(file)

	myLine, err := reader.ReadString('\n')
	check(err)

	/* Steps for Loading wad into game
	 * 1. Read a line
	 * 2. Determine its location (Screen vs Level)
	 * 3. Update Object
	 */
	for err != io.EOF {

		// Build out Map
		if strings.Contains(myLine, DELIMITER) {
			fmt.Print(".") // status load indicator
			
			newLevel := new(LevelDetails)
			newLevel.ParseAndHandleDetailsString(myLine)	

			myLine, err = reader.ReadString('\n')
			check(err)

			//Buid out the screen
			for idx := 0; idx < DISPLAYHEIGHT && err != io.EOF && !strings.Contains(myLine, DELIMITER); idx++ {
				newLevel.LoadRow(myLine, idx)
				myLine, err = reader.ReadString('\n')
				check(err)
			}

			game.AddLevel(newLevel.Id, newLevel)
		}

		myLine, err = reader.ReadString('\n')
	}

	fmt.Println("finished successfully")
}