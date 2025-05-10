package main

import (
	"github.com/ericharm/sogoban/domain"
	"github.com/ericharm/sogoban/models"
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

	game := domain.BuildLevel("data/2.lvl")
	defer models.CloseLogFile()

	for game.Running == true {
		game.Tick(stdscr)
	}
}
