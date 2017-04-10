package views

type CommandView struct {
	*View
}

func NewCommandView() *CommandView {
	return &CommandView{}
}

func (self *CommandView) Draw() {
	width, height = getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(nil, ColorGray2))

	// second line
	fillLine(0, height-1, NewColors(nil, ColorBackground))
}
