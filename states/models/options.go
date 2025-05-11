package states

import (
	"github.com/rthornton128/goncurses"
)

type Option struct {
	text   string
	action func()
	x      int
	y      int
}

func NewOption(text string, action func()) *Option {
	return &Option{
		text:   text,
		action: action,
		x:      0,
		y:      0,
	}
}

func (option *Option) SetPos(x, y int) {
	option.x = x
	option.y = y
}

type OptionsList struct {
	options             []*Option
	columnOffsets       []int
	selectedOptionIndex int
}

func NewOptionsList(options []*Option, columnOffsets []int, x int, y int) *OptionsList {
	columnIndex := 0
	xOffset := columnOffsets[columnIndex] + x
	yOffset := y

	for _, option := range options {
		option.SetPos(xOffset, yOffset)

		columnIndex = (columnIndex + 1) % len(columnOffsets)
		xOffset = columnOffsets[columnIndex] + x
		if columnIndex == 0 {
			yOffset += SelectLevelScreenYSpacing
		}
	}

	return &OptionsList{
		options:             options,
		columnOffsets:       columnOffsets,
		selectedOptionIndex: 0,
	}
}

func (list *OptionsList) HandleInput(key goncurses.Key) {
	verticalDistance := len(list.columnOffsets)
	switch key {
	case goncurses.KEY_UP:
		list.selectedOptionIndex -= verticalDistance
	case goncurses.KEY_DOWN:
		list.selectedOptionIndex += verticalDistance
	case goncurses.KEY_LEFT:
		list.selectedOptionIndex -= 1
	case goncurses.KEY_RIGHT:
		list.selectedOptionIndex += 1
	}

	maxOptionIndex := len(list.options) - 1
	if list.selectedOptionIndex < 0 {
		list.selectedOptionIndex = maxOptionIndex
	}
	if list.selectedOptionIndex > maxOptionIndex {
		list.selectedOptionIndex = 0
	}

	if key == goncurses.KEY_ENTER || key == goncurses.KEY_RETURN {
		selectedOption := list.options[list.selectedOptionIndex]
		selectedOption.action()
	}
}

func (list *OptionsList) Draw(window *goncurses.Window) {
	for _, option := range list.options {
		window.MovePrint(option.y, option.x, option.text)
	}

	selectedOption := list.options[list.selectedOptionIndex]
	window.Move(selectedOption.y, selectedOption.x-SelectedOptionCursorGutter)
}
