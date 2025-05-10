package util

import (
	"github.com/ericharm/gokoban/defs"
	"github.com/rthornton128/goncurses"
)

func InitCurses(window *goncurses.Window) {
	goncurses.Raw(true)   // turn on raw "uncooked" input
	goncurses.Echo(false) // turn echoing of typed characters off
	goncurses.Cursor(0)   // hide cursor
	window.Keypad(true)   // allow keypad input
	startColor()
}

func startColor() {
	goncurses.StartColor()
	goncurses.InitPair(defs.White, goncurses.C_WHITE, goncurses.C_BLACK)
	goncurses.InitPair(defs.Red, goncurses.C_RED, goncurses.C_BLACK)
	goncurses.InitPair(defs.Green, goncurses.C_GREEN, goncurses.C_BLACK)
	goncurses.InitPair(defs.Blue, goncurses.C_BLUE, goncurses.C_BLACK)
	goncurses.InitPair(defs.Yellow, goncurses.C_YELLOW, goncurses.C_BLACK)
	goncurses.InitPair(defs.Magenta, goncurses.C_MAGENTA, goncurses.C_BLACK)
	goncurses.InitPair(defs.Cyan, goncurses.C_CYAN, goncurses.C_BLACK)
}

func GetOffset(window *goncurses.Window, width, height int) (int, int) {
	maxX, maxY := window.MaxYX()
	offsetX := (maxX - width) / 2
	offsetY := (maxY - height) / 2
	return offsetX, offsetY
}
