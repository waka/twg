package views

import runewidth "github.com/mattn/go-runewidth"

type StatusView struct {
}

func NewStatusView() *StatusView {
	return &StatusView{}
}

func (view *StatusView) Draw(mode, status string) {
	width, height := getWindowSize()

	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	mode = "*" + mode + "*"
	drawText(mode, 1, height-2, NewColors(ForeGround(ColorPink), BackGround(ColorGray2)))

	x := width - runewidth.StringWidth(status) - 1
	drawText(status, x, height-2, NewColors(ForeGround(ColorBlue), BackGround(ColorGray2)))
}
