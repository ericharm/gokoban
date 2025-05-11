package entities

import "github.com/ericharm/gokoban/defs"

type Player struct {
	BaseEntity
	OnExit bool
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
		OnExit: false,
	}
}

func (player *Player) PushInDirection(direction Direction, entities map[defs.Vec2]Entity) bool {
	target := defs.Vec2{player.x + direction[0], player.y + direction[1]}
	other, exists := entities[target]

	if exists {
		ok := other.PushInDirection(direction, entities)
		if !ok {
			return false
		}
	}

	delete(entities, defs.Vec2{player.x, player.y})
	player.move(direction[0], direction[1])
	entities[defs.Vec2{player.x, player.y}] = player
	return true
}

func (player *Player) move(x, y int) {
	player.x += x
	player.y += y
}

func (player *Player) SetOnExit(onExit bool) {
	player.OnExit = onExit
}
