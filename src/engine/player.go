package engine

import (
	"fmt"
	"strings"
)

const(
	DEFAULTDIRECTIONICON = ""
)

var(
	//TODO: Hardcoded directions icon mapping; pull from file
	playerDirection = map[string] string {
		"north":"▲", // ▲ - Looking North
		"south":"▼", // ► - Looking East
		"east":"►",  // ▼ - Looking South
		"west":"◄",  // ◄ - Looking West
		"none":"◊",  // ◊ - Normal/Default
		"":"◊",      // ◊ - Normal/Default
	}
)

type PlayerDetails struct {
	Name   string
	Health int16
	Location *PositionDetails
	Icon string
}

func (p *PlayerDetails) Init(name string) {
	p.UpdateName(name)
	p.UpdateHealth(100)
	p.Location = new(PositionDetails)
	p.Icon = playerDirection[DEFAULTDIRECTIONICON]
}

func (p *PlayerDetails) UpdatePosition(x int, y int, direction string){
	p.Location.X = x
	p.Location.Y = y
	p.UpdatePlayerIcon(direction)
}

func (p *PlayerDetails) UpdatePlayerIcon(direction string){
	udirection := strings.ToLower(direction)
	if val, found := playerDirection[udirection]; found {
		p.Icon = val
		return
	}

	//TODO: log direction not found
	p.Icon = playerDirection[DEFAULTDIRECTIONICON]
}

func (p *PlayerDetails) UpdateName(val string) {
	p.Name = val
}

func (p *PlayerDetails) UpdateHealth(val int16) {
	p.Health += val
}

func (p *PlayerDetails) Print() {
	fmt.Println("[Player (", p.Name, ")] Health: ", p.Health, "/100")
}
