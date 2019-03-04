package engine

import (
	"fmt"
)

type PlayerDetails struct {
	Name   string
	Health int16
}

func (p PlayerDetails) Init(name string) {
	p.UpdateName(name)
	p.UpdateHealth(100)
}

func (p PlayerDetails) UpdateName(val string) {
	p.Name = val
}

func (p PlayerDetails) UpdateHealth(val int16) {
	p.Health += val
}

func (p PlayerDetails) Print() {
	fmt.Println("[Player (", p.Name, ")] Health: ", p.Health, "/100")
}
