package models

import "github.com/rthornton128/goncurses"
import "github.com/ericharm/sogoban/defs"

type Game struct {
	Player   *Player
	Entities map[defs.Point]Entity
}

func NewGame(player *Player, entities map[defs.Point]Entity) *Game {
	return &Game{
		Player:   player,
		Entities: entities,
	}
}

func (game *Game) Print(stdscr *goncurses.Window) {
	for _, entity := range game.Entities {
		entity.Print(stdscr)
	}

	game.Player.Print(stdscr)
}
