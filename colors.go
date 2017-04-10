package views

import (
	"github.com/nsf/termbox-go"
)

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

type colors struct {
	fg termbox.Attribute
	bg termbox.Attribute
}

func NewColors(fg termbox.Attribute, bg termbox.Attribute) *colors {
	return &colors{fg: fg, bg: bg}
}
