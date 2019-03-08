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
		if spaces < 0 {
			gd.MoveSouth(-1 * spaces)
		}
		gd.MoveNorth(spaces)
	case SOUTH:
		if spaces < 0 {
			gd.MoveNorth(-1 * spaces)
		}
		gd.MoveSouth(spaces)
	case EAST:
		if spaces < 0 {
			gd.MoveWest(-1 * spaces)
		}
		gd.MoveEast(spaces)
	case WEST:
		if spaces < 0 {
			gd.MoveEast(-1 * spaces)
		}
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
	spacesToMove, _ := gd.ObjectInPath(EAST, spaces)

	if gd.Player.Location.Y-spacesToMove < 0 {
		gd.Player.UpdatePosition(gd.Player.Location.X, 0, "north")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X, gd.Player.Location.Y-spacesToMove, "north")
	}
}

func (gd *GameDetails) MoveSouth(spaces int) {

	//TODO object detection
	spacesToMove, _ := gd.ObjectInPath(EAST, spaces)

	if gd.Player.Location.Y+spacesToMove >= DISPLAYHEIGHT {
		gd.Player.UpdatePosition(gd.Player.Location.X, DISPLAYHEIGHT-1, "south")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X, gd.Player.Location.Y+spacesToMove, "south")
	}
}

func (gd *GameDetails) MoveEast(spaces int) {

	//TODO object detection
	spacesToMove, _ := gd.ObjectInPath(EAST, spaces)

	if gd.Player.Location.X+spacesToMove >= DISPLAYWIDTH {
		gd.Player.UpdatePosition(DISPLAYWIDTH-1, gd.Player.Location.Y, "east")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X+spacesToMove, gd.Player.Location.Y, "east")
	}
}

func (gd *GameDetails) MoveWest(spaces int) {

	//TODO object detection
	spacesToMove, _ := gd.ObjectInPath(WEST, spaces)

	if gd.Player.Location.X-spacesToMove < 0 {
		gd.Player.UpdatePosition(0, gd.Player.Location.Y, "west")
	} else {
		gd.Player.UpdatePosition(gd.Player.Location.X-spacesToMove, gd.Player.Location.Y, "west")
	}
}

func (gd GameDetails) ObjectInPath(direction string, spaces int) (int, bool) {
	objectInPath := false
	spacesToMove := 0

	tileVector := gd.CurrentLevel.GetMoveTileVector(gd.Player.Location.X, gd.Player.Location.Y, direction, spaces)

	//fmt.Println(tileVector)

	for idx := 0; idx < len(tileVector); idx++ {
		if !tileVector[idx].IsMoveBlocker() {
			spacesToMove++
		} else {
			break
		}
	}

	return spacesToMove, objectInPath
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
