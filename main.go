package main

import (
	"github.com/ericharm/gokoban/states"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {
	stdscr, err := goncurses.Init()

	if err != nil {
		log.Fatal("init", err)
	}

	defer goncurses.End()

	util.InitCurses(stdscr)

	game := states.NewGameFromFile("data/2.lvl")
	defer util.CloseLogFile()

	for game.Running == true {
		game.Tick(stdscr)
	}
}
