package views

type CommandView struct {
	View
}

func NewCommandView() *CommandView {
	return &CommandView{}
}

func (self *CommandView) Draw() {
	_, height := getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	// second line
	fillLine(0, height-1, NewColors(BackGround(ColorBackground)))
}
