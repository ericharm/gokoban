package models

import "strconv"
import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"

type Game struct {
	Player   *Player
	Entities map[defs.Point]Entity
	turn     int
}

func NewGame(player *Player, entities map[defs.Point]Entity) *Game {
	return &Game{
		Player:   player,
		Entities: entities,
		turn:     0,
	}
}

func (game *Game) Print(stdscr *goncurses.Window) {
	for pt, entity := range game.Entities {
		x, y := entity.GetPos()
		if pt[0] != x {
			return
		}
		if pt[1] != y {

		}
		entity.Print(stdscr)
	}

	game.Player.Print(stdscr)
	game.turn += 1
}

func (game *Game) Log() {
	Log("Turn: " + strconv.Itoa(game.turn) + "\n")

	x, y := game.Player.GetPos()
	Log("Player: (")
	Log(strconv.Itoa(x))
	Log(", ")
	Log(strconv.Itoa(y))
	Log(")\n")

	for pt, entity := range game.Entities {
		x, y := entity.GetPos()
		Log(entity.GetChar())
		Log(" (")
		Log(strconv.Itoa(pt[0]))
		Log(", ")
		Log(strconv.Itoa(pt[1]))
		Log(") (")
		Log(strconv.Itoa(x))
		Log(", ")
		Log(strconv.Itoa(y))
		Log(")\n")
	}

	Log("\n")
}

func (game *Game) Close() {
	Close()
}
