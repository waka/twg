package views

import termbox "github.com/nsf/termbox-go"

type Container struct {
	viewMode     ViewMode
	timelineView *TimelineView
	mentionsView *MentionsView
	listView     *ListView
	commandView  *CommandView
}

type ViewMode int

const (
	MODE_TIMELINE ViewMode = iota + 1
	MODE_MENTION
	MODE_LIST
)

func NewContainer() *Container {
	return &Container{
		timelineView: NewTimelineView(),
		mentionsView: NewMentionsView(),
		listView:     NewListView(),
		commandView:  NewCommandView(),
	}
}

func (container *Container) Setup() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc | termbox.InputAlt)

	container.viewMode = MODE_TIMELINE

	return nil
}

func (container *Container) ChangeViewMode(viewMode ViewMode) {
	container.viewMode = viewMode
}

func (container *Container) StartRuneInCommand() {
}

func (container *Container) AddRuneInCommand(r rune) {
}

func (container *Container) ClearRuneInCommand() {
}

func (container *Container) Render() {
	container.RenderContents()
	container.RenderCommand()
}

func (container *Container) RenderContents() {
	termbox.Clear(ColorBackground, ColorBackground)

	switch container.viewMode {
	case MODE_TIMELINE:
		container.timelineView.Draw()
	case MODE_MENTION:
		container.mentionsView.Draw()
	case MODE_LIST:
		container.listView.Draw()
	}
	termbox.Flush()
}

func (container *Container) RenderCommand() {
	container.commandView.Draw(container.viewMode)
	termbox.Flush()
}

func (container *Container) Dispose() {
	termbox.Close()
}
