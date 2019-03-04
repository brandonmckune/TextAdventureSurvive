package engine

type GameDetails struct {
	Levels  map[string]*LevelDetails
	Screens map[string]*ScreenDetails
	Player  PlayerDetails
}

const (
	DISPLAYHEIGHT int = 40
	DISPLAYWIDTH  int = 80
)

func init() {

}

// func InitGame() *GameDetails {
// 	game := new(GameDetails)
// 	game.initLevels()
// 	game.initScreens()
// 	return game
// }

// func (gd *GameDetails) initLevels() {
// 	gd.Levels = make(map[string]*LevelDetails)
// }

func (gd *GameDetails) AddScreen(sequenceKey string, screen *ScreenDetails) {
	if gd.Screens == nil {
		gd.Screens = make(map[string]*ScreenDetails)
	}

	if _, found := gd.Screens[sequenceKey]; found {
		panicString := "Screen of sequence [" + sequenceKey + "] already exists!"
		panic(panicString)
	}

	gd.Screens[sequenceKey] = screen
}

func (gd *GameDetails) AddLevel(levelId string, level *LevelDetails) {
	if gd.Levels == nil {
		gd.Levels = make(map[string]*LevelDetails)
	}

	if _, found := gd.Levels[levelId]; found {
		panicString := "Level Id: [" + levelId + "] already exists!"
		panic(panicString)
	}

	gd.Levels[levelId] = level
}

func (gd *GameDetails) DisplayScreen(screenName string) {
	if screen, found := gd.Screens[screenName]; found {
		screen.Print()
	}
}

func (gd *GameDetails) DisplayLevel(levelName string) {
	if level, found := gd.Levels[levelName]; found {
		level.Print()
	}
}
