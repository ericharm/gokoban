package states

import (
	models "github.com/ericharm/gokoban/states/models"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
	"strconv"
)

type Game struct {
	level *models.Level
	turn  int
}

func NewGameFromFile(filename string) *Game {
	goncurses.Cursor(0)
	level := models.NewLevelFromFile(filename)
	return &Game{
		level: level,
		turn:  0,
	}
}

func (game *Game) Draw(window *goncurses.Window) {
	game.level.Draw(window)
	game.Log()
}

func (game *Game) HandleInput(char goncurses.Key) {
	game.level.HandleInput(char)
	game.turn += 1
}

func (game *Game) Log() {
	util.WriteToLog("Turn: " + strconv.Itoa(game.turn) + "\n")
	game.level.Log()
	util.WriteToLog("\n")
}
