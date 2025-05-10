package states

import (
	models "github.com/ericharm/gokoban/states/models"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
	"strconv"
)

type Game struct {
	Running bool
	level   *models.Level
	turn    int
}

func NewGameFromFile(filename string) *Game {
	level := models.NewLevelFromFile(filename)
	return &Game{
		Running: true,
		level:   level,
		turn:    0,
	}
}

func (game *Game) Tick(window *goncurses.Window) {
	game.turn += 1
	window.Clear()
	game.level.Print(window)
	window.Refresh()

	char := window.GetChar()
	if char == 'q' {
		game.Running = false
	} else {
		game.level.HandleInput(char)
	}

	game.Log()
}

func (game *Game) Log() {
	util.WriteToLog("Turn: " + strconv.Itoa(game.turn) + "\n")
	game.level.Log()
	util.WriteToLog("\n")
}
