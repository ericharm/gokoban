package main

import (
	"github.com/ericharm/sogoban/defs"
	"github.com/ericharm/sogoban/domain"
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {

	stdscr, err := goncurses.Init()
	if err != nil {
		log.Fatal("init", err)
	}
	defer goncurses.End()

	goncurses.Raw(true)   // turn on raw "uncooked" input
	goncurses.Echo(false) // turn echoing of typed characters off
	goncurses.Cursor(0)   // hide cursor
	stdscr.Keypad(true)   // allow keypad input
	domain.StartColor()

	game := domain.BuildLevel("data/1.lvl")
	defer game.Close()

	for {
		stdscr.Clear()
		game.Print(stdscr)
		stdscr.Refresh()

		char := stdscr.GetChar()
		switch char {
		case 'q':
			return
		case goncurses.KEY_UP:
			game.Player.PushInDirection(defs.Up, game.Entities)
		case goncurses.KEY_DOWN:
			game.Player.PushInDirection(defs.Down, game.Entities)
		case goncurses.KEY_LEFT:
			game.Player.PushInDirection(defs.Left, game.Entities)
		case goncurses.KEY_RIGHT:
			game.Player.PushInDirection(defs.Right, game.Entities)
		}

		game.Log()
	}
}
