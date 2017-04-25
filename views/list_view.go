package views

// ListView is one of views.
type ListView struct {
	View
}

// NewListView returns instance.
func NewListView() *ListView {
	return &ListView{}
}

// Draw ...
func (view *ListView) Draw() {
	_, height := getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	// second line
	fillLine(0, height-1, NewColors(BackGround(ColorBackground)))
}
