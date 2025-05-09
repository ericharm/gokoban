package defs

type Point [2]int

type EntityType int

const (
	EntityPlayer EntityType = iota
	EntityWall
	EntityBoulder
	EntityPit
	EntityExit
)
