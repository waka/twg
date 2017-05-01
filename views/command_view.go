package views

import (
	"unicode/utf8"

	runewidth "github.com/mattn/go-runewidth"
)

type CommandView struct {
	content []byte
	command []byte
	cursorX int
}

func NewCommandView() *CommandView {
	return &CommandView{cursorX: 0}
}

func (view *CommandView) Draw(viewMode ViewMode) {
	_, height := getWindowSize()

	// first line
	fillLine(0, height-2, NewColors(BackGround(ColorGray2)))

	mode := view.getViewModeString(viewMode)
	drawText(mode, 2, height-2, NewColors(ForeGround(ColorYellow), BackGround(ColorGray2)))

	// second line
	fillLine(0, height-1, NewColors(BackGround(ColorBackground)))
	view.drawCommandLine()
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

func (view *CommandView) handleEvent(event *CommandEvent) {
	switch event.eventType {
	case CommandStart:
		view.startCommand()
	case CommandEnd:
		view.clearCommand()
	case CommandAdd:
		view.addCommand(event.value)
	case CommandDelete:
		view.deleteCommand()
	case CommandLeft:
		view.moveLeftCommand()
	case CommandRight:
		view.moveRightCommand()
	case CommandExecute:
		view.executeCommand()
	}

	if event.eventType == CommandEnd {
		hideCursor()
	} else {
		view.focusCursor()
	}
}

func (view *CommandView) startCommand() {
	view.content = byteSliceInsert(view.content, []byte(":"), view.cursorX)
	view.moveCursorRight()
	view.drawCommandLine()
}

func (view *CommandView) clearCommand() {
	view.content = []byte{}
	view.cursorX = 0
	view.drawCommandLine()
}

func (view *CommandView) addCommand(command []byte) {
	view.content = byteSliceInsert(view.content, command, view.cursorX)
	view.moveCursorRight()
	view.drawCommandLine()
}

func (view *CommandView) deleteCommand() {
	if 2 > view.cursorX {
		return
	}
	view.moveCursorLeft()
	_, size := utf8.DecodeRune(view.content[view.cursorX:])
	view.content = byteSliceRemove(view.content, view.cursorX, view.cursorX+size)
	view.drawCommandLine()
}

func (view *CommandView) moveLeftCommand() {
	view.moveCursorLeft()
}

func (view *CommandView) moveRightCommand() {
	view.moveCursorRight()
}

func (view *CommandView) executeCommand() {}

func (view *CommandView) drawCommandLine() {
	_, height := getWindowSize()
	content := string(view.content)
	drawText(content, 2, height-1, NewColors(ForeGround(ColorWhite), BackGround(ColorBackground)))
}

func (view *CommandView) moveCursorLeft() {
	if view.cursorX <= 0 {
		return
	}
	_, size := utf8.DecodeLastRune(view.content[:view.cursorX])
	view.cursorX -= size
}

func (view *CommandView) moveCursorRight() {
	if view.cursorX >= len(view.content) {
		return
	}
	_, size := utf8.DecodeRune(view.content[view.cursorX:])
	view.cursorX += size
}

func (view *CommandView) focusCursor() {
	_, height := getWindowSize()
	x := runewidth.StringWidth(string(view.content[:view.cursorX]))
	x += 2
	setCursor(x, height-1)
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

func byteSliceRemove(bytes []byte, from int, to int) []byte {
	copy(bytes[from:], bytes[to:])
	return bytes[:len(bytes)+from-to]
}
