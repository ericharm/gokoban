package entities

import "github.com/ericharm/gokoban/defs"

type Wall struct {
	BaseEntity
}

func NewWall(x int, y int) Entity {
	return &Wall{
		BaseEntity: BaseEntity{
			entityType: EntityWall,
			x:          x,
			y:          y,
			char:       '#',
			color:      defs.White,
		},
	}
}
