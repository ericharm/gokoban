package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Entity struct {
	BaseNode
	X      int32
	Y      int32
	Radius float32
	Angle  float32
	Color  rl.Color
}

func NewEntity(x int32, y int32, radius float32, color rl.Color) *Entity {
	return &Entity{
		X:      x,
		Y:      y,
		Radius: radius,
		Angle:  0,
		Color:  color,
	}
}

func (e *Entity) Draw() {
	rl.DrawCircle(e.X, e.Y, e.Radius, e.Color)
}

func (e *Entity) Rotate(angle float32) {
	e.Angle += angle
	if e.Angle >= 360 {
		e.Angle -= 360
	}
}

func (e *Entity) Move(dx, dy int32) {
	e.X += dx
	e.Y += dy
}
