package states

import (
	models "github.com/ericharm/gokoban/states/models"
	"github.com/ericharm/gokoban/util"
	"github.com/rthornton128/goncurses"
)

type StageSelect struct {
	options *models.OptionsList
}

func NewStageSelect() *StageSelect {
	goncurses.Cursor(1)

	application := GetApplication()
	window := application.GetWindow()
	maxY, maxX := window.MaxYX()

	singleColumnWidth := len("Level 8") + models.SelectedOptionCursorGutter*2
	columnCount := models.SelectLevelScreenColumnCount
	optionsListWidth := singleColumnWidth * columnCount
	optionsListHeight := len(stageSelectOptions) / columnCount * models.SelectLevelScreenYSpacing

	centerX, centerY := util.GetOffset(maxX, maxY, optionsListWidth, optionsListHeight)

	return &StageSelect{
		options: models.NewOptionsList(
			stageSelectOptions, []int{0, singleColumnWidth}, centerX, centerY,
		),
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
