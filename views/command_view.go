package views

type CommandView struct {
	content []byte
	command []byte
}

func NewCommandView() *CommandView {
	return &CommandView{}
}

func (view *CommandView) Draw(viewMode ViewMode) {
	_, height := getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	mode := view.getViewModeString(viewMode)
	drawText(mode, 2, height-2, NewColors(ForeGround(ColorYellow), BackGround(ColorGray2)))

	// second line
	fillLine(0, height-1, NewColors(BackGround(ColorBackground)))

	content := string(view.content)
	drawText(content, 0, height-1, NewColors(ForeGround(ColorWhite), BackGround(ColorBackground)))
}

func (view *CommandView) getViewModeString(viewMode ViewMode) string {
	switch viewMode {
	case MODE_TIMELINE:
		return "*Timeline*"
	case MODE_MENTION:
		return "*Mention*"
	case MODE_LIST:
		return "*List*"
	}
	return "*No mode*"
}

func byteSliceRemove(bytes []byte, from int, to int) []byte {
	copy(bytes[from:], bytes[to:])
	return bytes[:len(bytes)+from-to]
}

func byteSliceInsert(dst []byte, src []byte, pos int) []byte {
	length := len(dst) + len(src)
	if cap(dst) < length {
		s := make([]byte, len(dst), length)
		copy(s, dst)
		dst = s
	}
	dst = dst[:length]
	copy(dst[pos+len(src):], dst[pos:])
	copy(dst[pos:], src)
	return dst

}
