package models

import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"

type Entity interface {
	GetType() defs.EntityType
	GetPos() (int, int)
	setPos(int, int)
	GetChar() string
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

func (entity *BaseEntity) GetChar() string {
	return string(entity.char)
}

func (entity *BaseEntity) Print(stdscr *goncurses.Window) {
	stdscr.MoveAddChar(
		entity.y,
		entity.x,
		entity.char|goncurses.ColorPair(entity.color),
	)
}
func (entity *BaseEntity) GetPos() (int, int) {
	return entity.x, entity.y

}

func (entity *BaseEntity) setPos(x, y int) {
	entity.x = x
	entity.y = y
}

func moveEntity(entity Entity, x, y int, entities map[defs.Point]Entity) {
	// remove the entity from its old position
	entityX, entityY := entity.GetPos()
	delete(entities, defs.Point{entityX, entityY})

	newX := entityX + x
	newY := entityY + y
	entity.setPos(newX, newY)

	// add the entity to its new position
	entities[defs.Point{newX, newY}] = entity
}

func (entity *BaseEntity) Move(x, y int, entities map[defs.Point]Entity) {
	moveEntity(entity, x, y, entities)
}

func (entity *BaseEntity) PushInDirection(direction defs.Direction, entities map[defs.Point]Entity) bool {
	return false
}
