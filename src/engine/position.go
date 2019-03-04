package engine

import(

)

const(

)

var(

)

type PositionDetails struct {
	X, Y int
}

func (p *PositionDetails) UpdatePosition(x int, y int){
	p.X = x
	p.Y = y
}