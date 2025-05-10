package domain

import (
	en "github.com/ericharm/sogoban/entities"
	"github.com/ericharm/sogoban/models"
	iofs "io/fs"
	"log"
	"os"
)

func NewEntityFromChar(ch string, x, y int) (en.Entity, *en.Player) {
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

func BuildLevel(filePath string) *models.Game {
	data, err := iofs.ReadFile(os.DirFS("."), filePath)

	if err != nil {
		log.Fatal("init", err)
	}

	player := en.NewPlayer(0, 0)
	entities := make(map[en.Point]en.Entity)

	x := 0
	y := 0
	for _, byte := range data {
		byteAsStr := string(byte)

		if byteAsStr == "\n" {
			x = 0
			y += 1
			continue
		}

		entity, p := NewEntityFromChar(byteAsStr, x, y)

		if p != nil {
			player = p
		}

		if entity != nil {
			entities[en.Point{x, y}] = entity
		}

		x += 1
	}

	return models.NewGame(player, entities)
}
