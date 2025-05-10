package domain

import (
	en "github.com/ericharm/gokoban/entities"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
	iofs "io/fs"
	"log"
	"os"
)

type Level struct {
	player   *en.Player
	entities map[en.Point]en.Entity
	width    int
	height   int
}

func (level *Level) HandleInput(char goncurses.Key) {
	switch char {
	case goncurses.KEY_UP:
		level.player.PushInDirection(en.Up, level.entities)
	case goncurses.KEY_DOWN:
		level.player.PushInDirection(en.Down, level.entities)
	case goncurses.KEY_LEFT:
		level.player.PushInDirection(en.Left, level.entities)
	case goncurses.KEY_RIGHT:
		level.player.PushInDirection(en.Right, level.entities)
	}
}

func (level *Level) Print(window *goncurses.Window) {
	offsetX, offsetY := util.GetOffset(window, level.width, level.height)
	offset := en.Point{offsetX, offsetY}

	for pt, entity := range level.entities {
		x, y := entity.GetPos()
		if pt[0] != x {
			return
		}
		if pt[1] != y {

		}
		entity.Print(window, offset)
	}

	level.player.Print(window, offset)
}

func (level *Level) Log() {
	util.WriteToLog(level.player.Debug())
	util.WriteToLog(" ")

	for _, entity := range level.entities {
		util.WriteToLog(entity.Debug())
		util.WriteToLog(" ")
	}
}

func NewLevelFromFile(filePath string) *Level {
	data, err := iofs.ReadFile(os.DirFS("."), filePath)

	if err != nil {
		log.Fatal("init", err)
	}

	player := en.NewPlayer(0, 0)
	entities := make(map[en.Point]en.Entity)

	width := 0
	x := 0
	y := 0
	for _, byte := range data {
		byteAsStr := string(byte)

		if byteAsStr == "\n" {
			width = max(width, x)
			x = 0
			y += 1
			continue
		}

		entity, p := newEntityFromChar(byteAsStr, x, y)

		if p != nil {
			player = p
		}

		if entity != nil {
			entities[en.Point{x, y}] = entity
		}

		x += 1
	}

	return &Level{player, entities, width, y}
}

func newEntityFromChar(ch string, x, y int) (en.Entity, *en.Player) {
	if ch == "@" {
		player := en.NewPlayer(x, y)
		return nil, player
	}

	var charToEntityCreateFuncMap = map[string]func(int, int) en.Entity{
		"#": en.NewWall,
		"0": en.NewBoulder,
		"^": en.NewPit,
		"X": en.NewExit,
	}

	creationFunc, ok := charToEntityCreateFuncMap[ch]

	if !ok {
		return nil, nil
	}

	return creationFunc(x, y), nil
}
