package main

import (
	"github.com/ericharm/gokoban/states"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
	"log"
)

func main() {
	window, err := goncurses.Init()

	if err != nil {
		log.Fatal("init", err)
	}

	defer goncurses.End()
	defer util.CloseLogFile()

	util.InitCurses(window)
	application := states.GetApplication()
	application.PushState(states.NewStageSelect(window))
	application.Run(window)
}
