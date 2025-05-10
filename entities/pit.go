package entities

import "github.com/ericharm/gokoban/defs"

type Pit struct {
	BaseEntity
}

func NewPit(x int, y int) Entity {
	return &Pit{
		BaseEntity: BaseEntity{
			entityType: EntityPit,
			x:          x,
			y:          y,
			char:       '^',
			color:      defs.Blue,
		},
	}
}
