package engine

import(
	"fmt"
)

const (
	NONPLAYERICON = " "
)

var(

)

type TileDetails struct {
	MapIcon string
	Location *PositionDetails
}

func (t *TileDetails) AddPlayer(p *PlayerDetails){
	t.MapIcon = p.Icon
}

func (t *TileDetails) RemovePlayer(){
	t.MapIcon = NONPLAYERICON
}

func (t *TileDetails) CheckLocationCreated() {
	if t.Location == nil {
		t.Location = new(PositionDetails)
	}
}

func (t *TileDetails) AddLocationPosition(pos *PositionDetails){
	t.CheckLocationCreated()
	t.Location = pos
}

func (t *TileDetails) AddLocationXY(x int, y int){
	t.CheckLocationCreated()
	t.Location.X = x
	t.Location.Y = y
}

func (t TileDetails) PrintTile(){
	fmt.Println("Tile: ", t.MapIcon, ", Position: ", t.Location.X, ", ", t.Location.Y)
}