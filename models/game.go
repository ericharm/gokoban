package models

import "strconv"
import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"
import "github.com/ericharm/sogoban/entities"

type Game struct {
	Player   *entities.Player
	Entities map[entities.Point]entities.Entity
	Running  bool
	turn     int
}

func NewGame(player *entities.Player, entities map[entities.Point]entities.Entity) *Game {
	return &Game{
		Player:   player,
		Entities: entities,
		Running:  true,
		turn:     0,
	}
}

func (game *Game) Tick(window *goncurses.Window) {
	game.turn += 1
	window.Clear()
	game.print(window)
	window.Refresh()

	char := window.GetChar()
	switch char {
	case 'q':
		game.Running = false
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

func (game *Game) print(window *goncurses.Window) {
	for pt, entity := range game.Entities {
		x, y := entity.GetPos()
		if pt[0] != x {
			return
		}
		if pt[1] != y {

		}
		entity.Print(window)
	}

	game.Player.Print(window)
}

func (game *Game) Log() {
	WriteToLog("Turn: " + strconv.Itoa(game.turn) + "\n")
	WriteToLog(game.Player.Debug())
	WriteToLog(" ")

	for _, entity := range game.Entities {
		WriteToLog(entity.Debug())
		WriteToLog(" ")
	}

	WriteToLog("\n")
}
