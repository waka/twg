package views

import (
	"github.com/nsf/termbox-go"
)

// Colors defs.
const (
	ColorBackground = termbox.ColorDefault
	ColorRed        = termbox.ColorRed
	ColorWhite      = termbox.ColorWhite
	ColorYellow     = termbox.ColorYellow
	ColorGreen      = termbox.ColorGreen
	ColorBlue       = termbox.ColorBlue
	ColorPink       = termbox.Attribute(214)
	ColorGray1      = termbox.Attribute(0xe9 + 9)
	ColorGray2      = termbox.Attribute(0xe9 + 6)
	ColorGray3      = termbox.Attribute(254)
	ColorLowlight   = termbox.Attribute(240)
	ColorBlack      = termbox.ColorBlack
)

type colorOptions struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

type Color func(*colorOptions)

func ForeGround(fg termbox.Attribute) Color {
	return func(co *colorOptions) {
		co.fg = fg
	}
}

func BackGround(bg termbox.Attribute) Color {
	return func(co *colorOptions) {
		co.bg = bg
	}
}

func NewColors(colors ...Color) *colorOptions {
	co := colorOptions{}
	for _, c := range colors {
		c(&co)
	}
	return &co
}
