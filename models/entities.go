package models

import "github.com/ericharm/sogoban/defs"

type Player struct {
	BaseEntity
}

func NewPlayer(x int, y int) *Player {
	return &Player{
		BaseEntity: BaseEntity{
			entityType: defs.EntityPlayer,
			x:          x,
			y:          y,
			char:       '@',
			color:      defs.Magenta,
		},
	}
}

func (player *Player) PushInDirection(direction defs.Direction, entities map[defs.Point]Entity) bool {
	target := defs.Point{player.x + direction[0], player.y + direction[1]}
	other, exists := entities[target]

	if exists {
		ok := other.PushInDirection(direction, entities)
		if !ok {
			return false
		}
	}

	player.Move(direction[0], direction[1], entities)
	return true
}

func (player *Player) Move(x, y int, entities map[defs.Point]Entity) {
	player.x += x
	player.y += y
}

type Wall struct {
	BaseEntity
}

func NewWall(x int, y int) Entity {
	return &Wall{
		BaseEntity: BaseEntity{
			entityType: defs.EntityWall,
			x:          x,
			y:          y,
			char:       '#',
			color:      defs.White,
		},
	}
}

type Boulder struct {
	BaseEntity
}

func NewBoulder(x int, y int) Entity {
	return &Boulder{
		BaseEntity: BaseEntity{
			entityType: defs.EntityBoulder,
			x:          x,
			y:          y,
			char:       '0',
			color:      defs.Cyan,
		},
	}
}

func (boulder *Boulder) PushInDirection(direction defs.Direction, entities map[defs.Point]Entity) bool {
	target := defs.Point{boulder.x + direction[0], boulder.y + direction[1]}
	other, exists := entities[target]

	if exists {
		if other.GetType() == defs.EntityPit {
			delete(entities, target)
			delete(entities, defs.Point{boulder.x, boulder.y})
			return true
		}
		return false
	}
	boulder.Move(direction[0], direction[1], entities)
	return true
}

func (entity *Boulder) Move(x, y int, entities map[defs.Point]Entity) {
	moveEntity(entity, x, y, entities)
}

type Pit struct {
	BaseEntity
}

func NewPit(x int, y int) Entity {
	return &Pit{
		BaseEntity: BaseEntity{
			entityType: defs.EntityPit,
			x:          x,
			y:          y,
			char:       '^',
			color:      defs.Blue,
		},
	}
}

type Exit struct {
	BaseEntity
}

func NewExit(x int, y int) Entity {
	return &Exit{
		BaseEntity: BaseEntity{
			entityType: defs.EntityExit,
			x:          x,
			y:          y,
			char:       'X',
			color:      defs.Green,
		},
	}
}
