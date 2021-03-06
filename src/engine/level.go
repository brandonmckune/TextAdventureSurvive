package engine

import (
	"fmt"
	"strconv"
	"strings"
)

type LevelDetails struct {
	Id                 string
	ConnectingLevels   map[string]*LevelDetails
	Details            [DISPLAYHEIGHT][DISPLAYWIDTH]*TileDetails
	StartPosition      *PositionDetails
	PlayerLastPosition *PositionDetails
}

var ()

const (
	STARTPOSITION = "startpos"
)

func init() {

}

func (d *LevelDetails) VerifyStartPosition() {
	if d.StartPosition == nil {
		d.StartPosition = new(PositionDetails)
	}
}

func (d *LevelDetails) VerifyPlayerLastPosition() {
	if d.PlayerLastPosition == nil {
		d.PlayerLastPosition = new(PositionDetails)
	}
}

/* LoadLevelRow allows the loader to read a single line from a file and insert that
 * data into the Level Details 2D array.
 */
func (d *LevelDetails) LoadRow(row string, rowIndex int) {
	if rowIndex < 0 || rowIndex > DISPLAYHEIGHT {
		panicString := "Invalid row index passed in LoadRow(" + string(rowIndex) + ")"
		fmt.Println(panicString)
		//panic(panicString)
	}

	//for idx, char := range row {
	for idx := 0; idx < len(row); idx++ {
		if idx >= DISPLAYWIDTH {
			fmt.Println("idx: ", idx, " Display: ", DISPLAYWIDTH)
			break
		}

		tile := new(TileDetails)
		tile.MapIcon = string(row[idx : idx+1])
		tile.AddLocationXY(rowIndex, idx)

		d.Details[rowIndex][idx] = tile
	}
}

func (d *LevelDetails) AddStartPosition(pos *PositionDetails) {
	d.StartPosition = pos
}

func (d *LevelDetails) AddStartPositionXY(x int, y int) {
	d.VerifyStartPosition()

	d.StartPosition.X = x
	d.StartPosition.Y = y
}

func (d *LevelDetails) AddConnectingLevel(key string, level *LevelDetails) {
	if _, found := d.ConnectingLevels[key]; !found {
		d.ConnectingLevels[key] = level
	}

	//TODO What do I do if it's there? Replace it?
}

func (d *LevelDetails) ParseAndHandleDetailsString(line string) {
	parts := strings.Split(line, FILE_DELIMITER)
	d.ParseAndHandleDetailsArray(parts)
}

func (d *LevelDetails) ParseAndHandleDetailsArray(parts []string) {

	if parts == nil || len(parts) <= 1 {
		panic("Invalid parsing for level details")
	}

	d.VerifyStartPosition()
	d.Id = parts[0] //The level Id is always the first and included Id

	if len(parts) == 1 { //We have nothing left to do, exit
		return
	}

	for idx := 1; idx < len(parts); idx++ {

		switch strings.ToLower(parts[idx]) {
		case STARTPOSITION:
			if idx+2 >= len(parts) {
				panic("Invalid level details string construction.")
			}

			x, _ := strconv.Atoi(parts[idx+1])
			y, _ := strconv.Atoi(parts[idx+2])
			d.AddStartPositionXY(x, y)
			idx = idx + 2
		}

	}
}

func (d *LevelDetails) AddPlayer(player *PlayerDetails) {
	d.VerifyPlayerLastPosition()
	d.PlayerLastPosition.X = player.Location.X
	d.PlayerLastPosition.Y = player.Location.Y
	d.Details[player.Location.Y][player.Location.X].AddPlayer(player)
}

func (d *LevelDetails) UpdatePlayer(player *PlayerDetails) {
	d.VerifyPlayerLastPosition()
	d.Details[d.PlayerLastPosition.Y][d.PlayerLastPosition.X].RemovePlayer()
	d.PlayerLastPosition.X = player.Location.X
	d.PlayerLastPosition.Y = player.Location.Y
	d.Details[player.Location.Y][player.Location.X].AddPlayer(player)
}

func (d *LevelDetails) GetMoveTileVector(startPosX int, startPosY int, direction string, spaces int) []TileDetails {
	tileVector := make([]TileDetails, spaces)
	tileVector = append(tileVector, *d.Details[startPosY][startPosX])

	switch direction {
	case NORTH:
		for i := 1; i <= spaces; i++ {
			if startPosY-i >= 0 {
				tileVector = append(tileVector, *d.Details[startPosY-i][startPosX])
			} else {
				i = spaces
			}
		}
	case SOUTH:
		for i := 1; i <= spaces; i++ {
			if startPosY+i < DISPLAYHEIGHT {
				tileVector = append(tileVector, *d.Details[startPosY+i][startPosX])
			} else {
				i = spaces
			}
		}
	case EAST:
		for i := 1; i <= spaces; i++ {
			if startPosX+i < DISPLAYWIDTH {
				tileVector = append(tileVector, *d.Details[startPosY][startPosX+i])
			} else {
				i = spaces
			}
		}
	case WEST:
		for i := 1; i <= spaces; i++ {
			if startPosX-i >= 0 {
				tileVector = append(tileVector, *d.Details[startPosY][startPosX-i])
			} else {
				i = spaces
			}
		}
	default:
		validDirections := "[" + NORTH + ", " + SOUTH + ", " + EAST + ", " + WEST + "]"
		fmt.Println("Invalid direction [", direction, "]. Values are ", validDirections)
		return nil
	}

	return tileVector
}

func (d LevelDetails) GetLevel(key string) *LevelDetails {
	if detail, found := d.ConnectingLevels[key]; found {
		return detail
	}

	panicString := "Level Not Found Exception was thrown!"
	fmt.Println(panicString)
	//panic(panicString)
	return nil
}

func (ld LevelDetails) Print() {
	for idy := 0; idy < len(ld.Details); idy++ {
		for idx := 0; idx < len(ld.Details[idy]); idx++ {
			fmt.Print(ld.Details[idy][idx].MapIcon)
		}
		fmt.Print("\n")
	}
}
