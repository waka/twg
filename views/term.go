package views

import (
	"github.com/mattn/go-runewidth"
	"github.com/nsf/termbox-go"
)

func getWindowSize() (int, int) {
	return termbox.Size()
}

func fillLine(offset int, y int, colors *colorOptions) {
	width, _ := getWindowSize()
	x := offset
	for {
		if x >= width {
			break
		}
		termbox.SetCell(x, y, rune(' '), ColorBackground, colors.bg)
		x++
	}
}

func drawText(str string, x int, y int, colors *colorOptions) {
	var fg, bg termbox.Attribute

	if colors != nil {
		fg = colors.fg
		bg = colors.bg
	} else {
		fg = ColorWhite
		bg = ColorBackground
	}
	i := 0
	for _, chr := range str {
		termbox.SetCell(x+i, y, rune(chr), fg, bg)
		i += runewidth.RuneWidth(chr)
	}
}

func setCursor(x int, y int) {
	termbox.SetCursor(x, y)
}

func hideCursor() {
	termbox.HideCursor()
}
