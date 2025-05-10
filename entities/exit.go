package entities

import "github.com/ericharm/gokoban/defs"

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
