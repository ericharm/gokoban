package entities

import "strconv"
import "strings"
import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"

type Entity interface {
	Print(*goncurses.Window)
	Debug() string
	PushInDirection(defs.Direction, map[Point]Entity) bool
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

func (entity *BaseEntity) Print(stdscr *goncurses.Window) {
	stdscr.MoveAddChar(
		entity.y,
		entity.x,
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

func (entity *BaseEntity) PushInDirection(direction defs.Direction, entities map[Point]Entity) bool {
	return false
}

func (entity *BaseEntity) GetPos() (int, int) {
	return entity.x, entity.y
}

func (entity *BaseEntity) GetEntityType() EntityType {
	return entity.entityType
}
