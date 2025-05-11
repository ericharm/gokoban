package states

import (
	models "github.com/ericharm/gokoban/states/models"
	"github.com/rthornton128/goncurses"
)

type StageSelect struct {
	options *models.OptionsList
}

func NewStageSelect(window *goncurses.Window) *StageSelect {
	goncurses.Cursor(1)
	return &StageSelect{
		options: models.NewOptionsList(stageSelectOptions, []int{-3, 8}, window, 3),
	}
}

func (s *StageSelect) Draw(window *goncurses.Window) {
	s.options.Draw(window)
}

func (s *StageSelect) HandleInput(key goncurses.Key) {
	s.options.HandleInput(key)
}

var stageSelectOptions = []*models.Option{
	models.NewOption("Level 1", func() { selectLevel("data/1.lvl") }),
	models.NewOption("Level 2", func() { selectLevel("data/2.lvl") }),
	models.NewOption("Level 3", func() { selectLevel("data/3.lvl") }),
	models.NewOption("Level 4", func() { selectLevel("data/4.lvl") }),
	models.NewOption("Level 5", func() { selectLevel("data/5.lvl") }),
	models.NewOption("Level 6", func() { selectLevel("data/6.lvl") }),
	models.NewOption("Level 7", func() { selectLevel("data/7.lvl") }),
	models.NewOption("Level 8", func() { selectLevel("data/8.lvl") }),
}

func selectLevel(fileName string) {
	application := GetApplication()
	application.SwapState(NewGameFromFile(fileName))
}
