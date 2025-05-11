package states

import "github.com/rthornton128/goncurses"

type State interface {
	Draw(window *goncurses.Window)
	HandleInput(goncurses.Key)
}
