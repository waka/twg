package views

import (
	"math/rand"

	"github.com/nsf/termbox-go"
)

// Color cache
var (
	PresetColors = []termbox.Attribute{
		ColorRed,
		ColorGreen,
		termbox.Attribute(40),
		termbox.Attribute(41),
		termbox.Attribute(64),
		termbox.Attribute(66),
		termbox.Attribute(70),
		termbox.Attribute(100),
		termbox.Attribute(110),
		termbox.Attribute(119),
		termbox.Attribute(124),
		termbox.Attribute(126),
		termbox.Attribute(130),
		termbox.Attribute(150),
		termbox.Attribute(160),
		termbox.Attribute(167),
		termbox.Attribute(216),
		termbox.Attribute(227),
	}
	ColorMap = make(map[int64]int, 50)
)

// Colors defs.
const (
	ColorBackground = termbox.ColorDefault
	ColorWhite      = termbox.ColorWhite
	ColorRed        = termbox.Attribute(204)
	ColorYellow     = termbox.Attribute(4)
	ColorGreen      = termbox.Attribute(107)
	ColorBlue       = termbox.Attribute(70)
	ColorPink       = termbox.Attribute(214)
	ColorGray1      = termbox.Attribute(246)
	ColorGray2      = termbox.Attribute(0xe9 + 6)
	ColorHighlight  = termbox.Attribute(244)
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

func GetRandomColor(id int64) termbox.Attribute {
	if val, ok := ColorMap[id]; ok {
		return PresetColors[val]
	}

	rand.Seed(id)
	val := rand.Intn(len(PresetColors))
	ColorMap[id] = val
	return PresetColors[val]
}
