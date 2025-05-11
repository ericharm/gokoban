package entities

import "strconv"
import "strings"
import "github.com/rthornton128/goncurses"

type Entity interface {
	Print(*goncurses.Window, Point)
	Debug() string
	PushInDirection(Direction, map[Point]Entity) bool
	GetPos() (int, int)
	GetEntityType() EntityType
}

type BaseEntity struct {
	entityType EntityType
	x          int
	y          int
	char       goncurses.Char
	color      int16
}

func (entity *BaseEntity) Print(window *goncurses.Window, offset Point) {
	window.MoveAddChar(
		entity.y+offset[1],
		entity.x+offset[0],
		entity.char|goncurses.ColorPair(entity.color),
	)
}

func (entity *BaseEntity) Debug() string {
	var builder strings.Builder
	builder.WriteString(string(entity.char))
	builder.WriteString(": (")
	builder.WriteString(strconv.Itoa(entity.x))
	builder.WriteString(", ")
	builder.WriteString(strconv.Itoa(entity.y))
	builder.WriteString(")")
	return builder.String()
}

func (entity *BaseEntity) PushInDirection(direction Direction, entities map[Point]Entity) bool {
	return false
}

func (entity *BaseEntity) GetPos() (int, int) {
	return entity.x, entity.y
}

func (entity *BaseEntity) GetEntityType() EntityType {
	return entity.entityType
}
