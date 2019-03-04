package engine

import(
	"fmt"
)

const (
	DISPLAYHEIGHT int = 40
	DISPLAYWIDTH  int = 80
	GAMESTART string = "0"
)

type GameDetails struct {
	Levels  map[string]*LevelDetails
	Screens map[string]*ScreenDetails
	LastScreen *ScreenDetails
	Player  PlayerDetails
}

func init() {

}

func (gd *GameDetails) NewGame(){
	gd.DisplayLevel(GAMESTART)
}

func (gd *GameDetails) validateScreens(){
	if gd.Screens == nil {
		gd.Screens = make(map[string]*ScreenDetails)
	}
}

func (gd *GameDetails) validateLevels(){
	if gd.Levels == nil {
		gd.Levels = make(map[string]*LevelDetails)
	}
}

func (gd *GameDetails) AddScreen(sequenceKey string, screen *ScreenDetails) {
	gd.validateScreens()

	if _, found := gd.Screens[sequenceKey]; found {
		panicString := "Screen of sequence [" + sequenceKey + "] already exists!"
		panic(panicString)
	}

	gd.Screens[sequenceKey] = screen
}

func (gd *GameDetails) AddLevel(levelId string, level *LevelDetails) {
	gd.validateLevels()

	if _, found := gd.Levels[levelId]; found {
		panicString := "Level Id: [" + levelId + "] already exists!"
		panic(panicString)
	}

	gd.Levels[levelId] = level
}

func (gd *GameDetails) DisplayScreen(screenName string) {
	gd.validateScreens()

	if screen, found := gd.Screens[screenName]; found {
		ClearConsole()
		screen.Print()
	} else {
		//TODO: Log invalid display
		fmt.Println("Screen is invalid! Screen Name:[", screenName, "]")
	}
}

func (gd *GameDetails) DisplayLevel(levelName string) {
	gd.validateLevels()

	if level, found := gd.Levels[levelName]; found {
		ClearConsole()
		level.Print()
	} else {
		//TODO: Log invalid display
		fmt.Println("Level is invalid! Level Name:[", levelName, "]")
	}
}
