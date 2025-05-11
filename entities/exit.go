package entities

import (
	"github.com/ericharm/gokoban/defs"
	"log"
)

type Exit struct {
	BaseEntity
}

func NewExit(x int, y int) Entity {
	return &Exit{
		BaseEntity: BaseEntity{
			entityType: EntityExit,
			x:          x,
			y:          y,
			char:       'X',
			color:      defs.Green,
		},
	}
}

func (exit *Exit) PushInDirection(direction Direction, entities map[Point]Entity) bool {
	pushedFrom := Point{exit.x - direction[0], exit.y - direction[1]}
	pusher, exists := entities[pushedFrom]

	if exists {
		if pusher.GetEntityType() == EntityPlayer {
			player, ok := pusher.(*Player)
			if !ok {
				log.Fatalf("Expected Player type, got %T", pusher)
			}
			player.SetOnExit(true)
		}
	}

	return false
}
