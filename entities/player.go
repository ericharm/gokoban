package entities

import "github.com/ericharm/gokoban/defs"

type Player struct {
	BaseEntity
}

func NewPlayer(x int, y int) *Player {
	return &Player{
		BaseEntity: BaseEntity{
			entityType: EntityPlayer,
			x:          x,
			y:          y,
			char:       '@',
			color:      defs.Magenta,
		},
	}
}

func (player *Player) PushInDirection(direction Direction, entities map[Point]Entity) bool {
	target := Point{player.x + direction[0], player.y + direction[1]}
	other, exists := entities[target]

	if exists {
		ok := other.PushInDirection(direction, entities)
		if !ok {
			return false
		}
	}

	player.move(direction[0], direction[1])
	return true
}

func (player *Player) move(x, y int) {
	player.x += x
	player.y += y
}
