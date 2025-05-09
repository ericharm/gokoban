package defs

type Direction []int

var (
	Left  Direction = []int{-1, 0}
	Right Direction = []int{1, 0}
	Up    Direction = []int{0, -1}
	Down  Direction = []int{0, 1}
)
