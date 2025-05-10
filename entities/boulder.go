package entities

import "github.com/ericharm/sogoban/defs"

type Boulder struct {
	BaseEntity
}

func NewBoulder(x int, y int) Entity {
	return &Boulder{
		BaseEntity: BaseEntity{
			entityType: EntityBoulder,
			x:          x,
			y:          y,
			char:       '0',
			color:      defs.Cyan,
		},
	}
}

func (boulder *Boulder) PushInDirection(direction defs.Direction, entities map[Point]Entity) bool {
	target := Point{boulder.x + direction[0], boulder.y + direction[1]}
	other, exists := entities[target]

	if exists {
		if other.GetEntityType() == EntityPit {
			delete(entities, target)
			delete(entities, Point{boulder.x, boulder.y})
			return true
		}
		return false
	}

	boulder.move(direction[0], direction[1], entities)
	return true
}

func (boulder *Boulder) move(x, y int, entities map[Point]Entity) {
	// remove the boulder from its old position
	boulderX, boulderY := boulder.GetPos()
	delete(entities, Point{boulderX, boulderY})

	newX := boulderX + x
	newY := boulderY + y
	boulder.x = newX
	boulder.y = newY

	// add the boulder to its new position
	entities[Point{newX, newY}] = boulder
}
