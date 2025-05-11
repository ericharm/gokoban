package entities

import "github.com/ericharm/gokoban/defs"

type EntityType int

const (
	EntityPlayer EntityType = iota
	EntityWall
	EntityBoulder
	EntityPit
	EntityExit
)

type Direction defs.Vec2

var (
	Left  Direction = [2]int{-1, 0}
	Right Direction = [2]int{1, 0}
	Up    Direction = [2]int{0, -1}
	Down  Direction = [2]int{0, 1}
)

func (d Direction) AsTuple() (int, int) {
	return d[0], d[1]
}

func (d Direction) AsVec2() defs.Vec2 {
	return defs.Vec2{d[0], d[1]}
}
