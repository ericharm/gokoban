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

func (o *Option) SetPos(x, y int) {
	o.x = x
	o.y = y
}

type OptionsList struct {
	options             []*Option
	columnOffsets       []int
	selectedOptionIndex int
}

func NewOptionsList(
	options []*Option,
	columnOffsets []int,
	startX int,
	startY int,
) *OptionsList {
	columnIndex := 0
	xOffset := columnOffsets[columnIndex] + startX
	yOffset := startY

	for _, option := range options {
		option.SetPos(xOffset, yOffset)

		columnIndex = (columnIndex + 1) % len(columnOffsets)
		xOffset = columnOffsets[columnIndex] + startX
		if columnIndex == 0 {
			yOffset += YSpacing
		}
	}

	return &OptionsList{
		options:             options,
		columnOffsets:       columnOffsets,
		selectedOptionIndex: 0,
	}
}

func (o *OptionsList) HandleInput(key goncurses.Key) {
	verticalDistance := len(o.columnOffsets)
	switch key {
	case goncurses.KEY_UP:
		o.selectedOptionIndex -= verticalDistance
	case goncurses.KEY_DOWN:
		o.selectedOptionIndex += verticalDistance
	case goncurses.KEY_LEFT:
		o.selectedOptionIndex -= 1
	case goncurses.KEY_RIGHT:
		o.selectedOptionIndex += 1
	}

	maxOptionIndex := len(o.options) - 1
	if o.selectedOptionIndex < 0 {
		o.selectedOptionIndex = maxOptionIndex
	}
	if o.selectedOptionIndex > maxOptionIndex {
		o.selectedOptionIndex = 0
	}

	if key == goncurses.KEY_ENTER || key == goncurses.KEY_RETURN {
		selectedOption := o.options[o.selectedOptionIndex]
		selectedOption.action()
	}
}

func (o *OptionsList) Draw(window *goncurses.Window) {
	for _, option := range o.options {
		window.MovePrint(option.y, option.x, option.text)
	}

	o.drawCursor(window)
}

func (o *OptionsList) drawCursor(window *goncurses.Window) {
	selectedOption := o.options[o.selectedOptionIndex]
	window.Move(selectedOption.y, selectedOption.x-SelectedOptionCursorGutter)
}
