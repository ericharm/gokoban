package domain

import (
	"github.com/ericharm/sogoban/defs"
	"github.com/ericharm/sogoban/models"
	iofs "io/fs"
	"log"
	"os"
)

func NewEntityFromChar(ch string, x, y int) (models.Entity, *models.Player) {

	if ch == "@" {
		player := models.NewPlayer(x, y)
		return nil, player
	}

	var charToEntityCreateFuncMap = map[string]func(int, int) models.Entity{
		"#": models.NewWall,
		"0": models.NewBoulder,
		"^": models.NewPit,
		"X": models.NewExit,
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

	player := models.NewPlayer(0, 0)
	entities := make(map[defs.Point]models.Entity)

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
			entities[defs.Point{x, y}] = entity
		}

		x += 1
	}

	return models.NewGame(player, entities)
}
