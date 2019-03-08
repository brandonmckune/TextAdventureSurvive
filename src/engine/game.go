package engine

import (
	"fmt"
	"strings"
)

const (
	DISPLAYHEIGHT int    = 40
	DISPLAYWIDTH  int    = 80
	GAMESTART     string = "0"
	NORTH         string = "NORTH"
	SOUTH         string = "SOUTH"
	EAST          string = "EAST"
	WEST          string = "WEST"
)

type GameDetails struct {
	Levels       map[string]*LevelDetails
	Screens      map[string]*ScreenDetails
	LastScreen   *ScreenDetails
	CurrentLevel *LevelDetails
	Player       *PlayerDetails
}

func init() {

}

func (gd *GameDetails) NewGame() {
	gd.Player = new(PlayerDetails)
	gd.Player.Init("Jeremy")
	gd.CurrentLevel = gd.GetLevel(GAMESTART)

	gd.Player.Location = gd.CurrentLevel.StartPosition
	gd.CurrentLevel.AddPlayer(gd.Player)

	gd.DisplayLevel()
}

func (gd *GameDetails) Move(direction string, spaces int) *GameDetails {
	//TODO object detection
	//Adjust spaces based on object detection

	successful := true
	direction = strings.ToUpper(direction)

	switch direction {
	case NORTH:
		gd.MoveNorth(spaces)
	case SOUTH:
		gd.MoveSouth(spaces)
	case EAST:
		gd.MoveEast(spaces)
	case WEST:
		gd.MoveWest(spaces)
	default:
		successful = false
		validDirections := "[" + NORTH + ", " + SOUTH + ", " + EAST + ", " + WEST + "]"
		fmt.Println("Invalid direction [", direction, "]. Values are ", validDirections)
	}

	if successful {
		gd.CurrentLevel.UpdatePlayer(gd.Player)
		gd.DisplayLevel()
	}

	return gd
}

func (gd *GameDetails) MoveNorth(spaces int) {

	//TODO object detection

	if gd.Player.Location.X-spaces < 0 {
		gd.Player.UpdatePosition(0, gd.Player.Location.Y, "north")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X-spaces, gd.Player.Location.Y, "north")
	}
}

func (gd *GameDetails) MoveSouth(spaces int) {

	//TODO object detection

	if gd.Player.Location.X+spaces >= DISPLAYHEIGHT {
		gd.Player.UpdatePosition(DISPLAYHEIGHT-1, gd.Player.Location.Y, "south")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X+spaces, gd.Player.Location.Y, "south")
	}
}

func (gd *GameDetails) MoveEast(spaces int) {

	//TODO object detection

	if gd.Player.Location.Y+spaces >= DISPLAYWIDTH {
		gd.Player.UpdatePosition(gd.Player.Location.X, DISPLAYWIDTH-1, "east")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X, gd.Player.Location.Y+spaces, "east")
	}
}

func (gd *GameDetails) MoveWest(spaces int) {

	//TODO object detection

	if gd.Player.Location.Y-spaces < 0 {
		gd.Player.UpdatePosition(gd.Player.Location.X, 0, "west")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X, gd.Player.Location.Y-spaces, "west")
	}
}

func (gd *GameDetails) validateScreens() {
	if gd.Screens == nil {
		gd.Screens = make(map[string]*ScreenDetails)
	}
}

func (gd *GameDetails) validateLevels() {
	if gd.Levels == nil {
		gd.Levels = make(map[string]*LevelDetails)
	}
}

func (gd *GameDetails) AddScreen(sequenceKey string, screen *ScreenDetails) {
	gd.validateScreens()

	if _, found := gd.Screens[sequenceKey]; found {
		panicString := "Screen of sequence [" + sequenceKey + "] already exists!"
		//fmt.Println(panicString)
		panic(panicString)
	}

	gd.Screens[sequenceKey] = screen
}

func (gd *GameDetails) AddLevel(levelId string, level *LevelDetails) {
	gd.validateLevels()

	if _, found := gd.Levels[levelId]; found {
		panicString := "Level Id: [" + levelId + "] already exists!"
		//fmt.Println(panicString)
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

func (gd *GameDetails) GetLevel(levelName string) *LevelDetails {
	gd.validateLevels()

	if level, found := gd.Levels[levelName]; found {
		return level
	}

	//TODO: Log invalid display
	panicString := "Level is invalid! Level Name:[" + levelName + "]"
	//fmt.Println(panicString)
	panic(panicString)
}

func (gd *GameDetails) DisplayLevel() {
	ClearConsole()
	gd.CurrentLevel.Print()
}

func (gd *GameDetails) DisplayLevelByName(levelName string) {
	gd.validateLevels()

	if level, found := gd.Levels[levelName]; found {
		ClearConsole()
		level.Print()
	} else {
		//TODO: Log invalid display
		fmt.Println("Level is invalid! Level Name:[", levelName, "]")
	}
}
