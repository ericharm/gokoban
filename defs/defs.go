package defs

const White = 0
const Red = 1
const Green = 2
const Blue = 3
const Yellow = 4
const Magenta = 5
const Cyan = 6

type Vec2 [2]int

func (v Vec2) AsTuple() (int, int) {
	return v[0], v[1]
}
