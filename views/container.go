package views

import (
	termbox "github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
)

type Container struct {
	tweetsView  *TweetsView
	statusView  *StatusView
	commandView *CommandView
}

func NewContainer() *Container {
	return &Container{
		tweetsView:  NewTweetsView(),
		statusView:  NewStatusView(),
		commandView: NewCommandView(),
	}
}

func (container *Container) Setup() error {
	if err := termbox.Init(); err != nil {
		return err
	}
	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputEsc | termbox.InputAlt)

	container.Clear()

	GetCommandEventEmitter().AddEventListener(container.commandView.handleEvent)

	return nil
}

func (container *Container) GetCommandValue() []byte {
	return container.commandView.GetValue()
}

func (container *Container) DrawTweets(tweets []*twitter.Tweet) {
	container.tweetsView.Draw(tweets)
}

func (container *Container) DrawStatus(mode, status string) {
	container.statusView.Draw(mode, status)
}

func (container *Container) DrawCommand() {
	container.commandView.Draw()
}

func (container *Container) Clear() {
	termbox.Clear(ColorBackground, ColorBackground)
}

func (container *Container) Render() {
	termbox.Flush()
}

func (container *Container) UpSelectedTweet(tweets []*twitter.Tweet) {
	container.tweetsView.UpCursor(tweets, container.tweetsView.cursorY-1)
}

func (container *Container) DownSelectedTweet(tweets []*twitter.Tweet) {
	container.tweetsView.DownCursor(tweets, container.tweetsView.cursorY+1)
}

func (container *Container) MoveToTopSelectedTweet(tweets []*twitter.Tweet) {
	container.tweetsView.UpCursor(tweets, 0)
}

func (container *Container) MoveToBottomSelectedTweet(tweets []*twitter.Tweet) {
	container.tweetsView.DownCursor(tweets, len(tweets)-1)
}

func (container *Container) GetSelectedTweet(tweets []*twitter.Tweet) *twitter.Tweet {
	return tweets[container.tweetsView.GetCursorPosition()]
}

func (container *Container) Dispose() {
	GetCommandEventEmitter().RemoveEventListener(container.commandView.handleEvent)
	termbox.Close()
}
