package models

import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"

type Entity interface {
	GetType() defs.EntityType
	Print(*goncurses.Window)
	Move(int, int, map[defs.Point]Entity)
	PushInDirection(defs.Direction, map[defs.Point]Entity) bool
}

type BaseEntity struct {
	entityType defs.EntityType
	x          int
	y          int
	char       goncurses.Char
	color      int16
}

func (entity *BaseEntity) GetType() defs.EntityType {
	return entity.entityType
}

func (entity *BaseEntity) Print(stdscr *goncurses.Window) {
	stdscr.MoveAddChar(
		entity.y,
		entity.x,
		entity.char|goncurses.ColorPair(entity.color),
	)
}

func (entity *BaseEntity) Move(x, y int, entities map[defs.Point]Entity) {
	// remove the entity from its old position
	delete(entities, [2]int{entity.x, entity.y})

	entity.x += x
	entity.y += y

	// add the entity to its new position
	entities[[2]int{entity.x, entity.y}] = entity
}

func (entity *BaseEntity) PushInDirection(direction defs.Direction, entities map[defs.Point]Entity) bool {
	target := [2]int{entity.x + direction[0], entity.y + direction[1]}
	_, exists := entities[target]

	if exists {
		return false
	}

	entity.Move(direction[0], direction[1], entities)
	return true
}

func NewEntity(entityType defs.EntityType, x int, y int) *BaseEntity {
	entity := &BaseEntity{
		entityType: entityType,
		x:          x,
		y:          y,
	}

	return entity
}

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
