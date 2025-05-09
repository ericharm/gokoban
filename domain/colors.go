package domain

import (
	"github.com/ericharm/sogoban/defs"
	"github.com/rthornton128/goncurses"
)

func StartColor() {
	goncurses.StartColor()
	goncurses.InitPair(defs.White, goncurses.C_WHITE, goncurses.C_BLACK)
	goncurses.InitPair(defs.Red, goncurses.C_RED, goncurses.C_BLACK)
	goncurses.InitPair(defs.Green, goncurses.C_GREEN, goncurses.C_BLACK)
	goncurses.InitPair(defs.Blue, goncurses.C_BLUE, goncurses.C_BLACK)
	goncurses.InitPair(defs.Yellow, goncurses.C_YELLOW, goncurses.C_BLACK)
	goncurses.InitPair(defs.Magenta, goncurses.C_MAGENTA, goncurses.C_BLACK)
	goncurses.InitPair(defs.Cyan, goncurses.C_CYAN, goncurses.C_BLACK)
}
