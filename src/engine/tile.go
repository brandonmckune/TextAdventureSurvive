package engine

import (
	"fmt"
)

const (
	NONPLAYERICON = " "
)

var ()

type TileDetails struct {
	OldIcon         string
	MapIcon         string
	Location        *PositionDetails
	IsVisualBlocker bool
}

func (self TileDetails) IsMoveBlocker() bool {
	if self.MapIcon == NONPLAYERICON {
		return false
	}

	for _, val := range playerDirection {
		if self.MapIcon == val {
			return false
		}
	}

	return true
}

func (t *TileDetails) AddPlayer(p *PlayerDetails) {
	t.OldIcon = t.MapIcon
	t.MapIcon = p.Icon
}

func (t *TileDetails) RemovePlayer() {
	t.MapIcon = t.OldIcon
}

func (t *TileDetails) CheckLocationCreated() {
	if t.Location == nil {
		t.Location = new(PositionDetails)
	}
}

func (t *TileDetails) AddLocationPosition(pos *PositionDetails) {
	t.CheckLocationCreated()
	t.Location = pos
}

func (t *TileDetails) AddLocationXY(x int, y int) {
	t.CheckLocationCreated()
	t.Location.X = x
	t.Location.Y = y
}

func (t TileDetails) PrintTile() {
	fmt.Println("Tile: ", t.MapIcon, ", Position: ", t.Location.X, ", ", t.Location.Y)
}
