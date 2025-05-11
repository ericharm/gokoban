package entities

import "github.com/ericharm/gokoban/defs"

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

func (boulder *Boulder) PushInDirection(direction Direction, entities map[defs.Vec2]Entity) bool {
	target := defs.Vec2{boulder.x + direction[0], boulder.y + direction[1]}
	other, exists := entities[target]

	if exists {
		if other.GetEntityType() == EntityPit {
			delete(entities, target)
			delete(entities, defs.Vec2{boulder.x, boulder.y})
			return true
		}
		return false
	}

	boulder.move(direction.AsVec2(), entities)
	return true
}

func (boulder *Boulder) move(by defs.Vec2, entities map[defs.Vec2]Entity) {
	// remove the boulder from its old position
	boulderX, boulderY := boulder.GetPos().AsTuple()
	delete(entities, defs.Vec2{boulderX, boulderY})

	x, y := by.AsTuple()
	newX := boulderX + x
	newY := boulderY + y
	boulder.x = newX
	boulder.y = newY

	// add the boulder to its new position
	entities[defs.Vec2{newX, newY}] = boulder
}
