package views

type ListView struct {
	View
}

func NewListView() *ListView {
	return &ListView{}
}

func (self *ListView) Draw() {
	_, height := getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	// second line
	fillLine(0, height-1, NewColors(BackGround(ColorBackground)))
}
