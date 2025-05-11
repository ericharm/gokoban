package util

import (
	"github.com/ericharm/gokoban/defs"
	"github.com/rthornton128/goncurses"
)

func InitCurses(window *goncurses.Window) {
	goncurses.Raw(true)
	goncurses.Echo(false)
	goncurses.Cursor(0)
	window.Keypad(true)
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

func GetOffset(maxX, maxY, width, height int) (int, int) {
	offsetX := (maxX - width) / 2
	offsetY := (maxY - height) / 2
	return offsetX, offsetY
}
