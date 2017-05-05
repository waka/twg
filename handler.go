package main

import (
	"fmt"
	"time"
	"unicode/utf8"

	termbox "github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
	"github.com/waka/twg/views"
)

// EventHandler
type Handler struct {
	command     *Command
	consumer    *Consumer
	container   *views.Container
	commandMode bool
	quit        bool
}

func NewHandler(command *Command, consumer *Consumer) *Handler {
	return &Handler{
		command:     command,
		consumer:    consumer,
		commandMode: false,
		quit:        false,
	}
}

func (handler *Handler) MainLoop() error {
	if err := handler.initialize(); err != nil {
		return err
	}
	defer handler.finish()

	eventCh := make(chan termbox.Event)
	defer close(eventCh)
	go func() {
		for {
			eventCh <- termbox.PollEvent()
		}
	}()

	timer := time.NewTicker(90 * time.Second)
	defer timer.Stop()

	for {
		select {
		case event := <-eventCh:
			if event.Type == termbox.EventResize {
				handler.reset()
			} else {
				switch GetActionTypeByEvent(event) {
				case ACTION_TYPE_ALL:
					handler.handleKeyEvent(event)
				case ACTION_TYPE_BOTH:
					handler.handleBothKeyEvent(event)
				case ACTION_TYPE_TWEET:
					handler.handleTweetKeyEvent(event)
				default:
					handler.handleCommandKeyEvent(event)
				}
			}
		case <-timer.C:
			if handler.command.IsReloadable() {
				handler.reload()
			}
		}
		if handler.quit {
			break
		}
	}

	return nil
}

func (handler *Handler) initialize() error {
	if err := handler.setupContainer(); err != nil {
		return err
	}

	// tweets for first view
	if err := handler.loadTweets(); err != nil {
		return err
	}
	handler.reset()

	return nil
}

func (handler *Handler) setupContainer() error {
	handler.container = views.NewContainer()
	if err := handler.container.Setup(); err != nil {
		return err
	}
	return nil
}

func (handler *Handler) handleKeyEvent(event termbox.Event) {
	if GetActionByEvent(event) == ACTION_QUIT {
		handler.quit = true
		return
	}

	eventEmitter := views.GetCommandEventEmitter()

	switch GetActionByEvent(event) {
	case ENTER_NORMAL_MODE:
		handler.commandMode = false
		eventEmitter.Emit(views.CommandEnd)
	case ENTER_COMMAND_MODE:
		handler.commandMode = true
		eventEmitter.Emit(views.CommandStart)
	}

	handler.reset()
}

func (handler *Handler) handleBothKeyEvent(event termbox.Event) {
	if handler.commandMode {
		handler.handleCommandKeyEvent(event)
	} else {
		handler.handleTweetKeyEvent(event)
	}
}

func (handler *Handler) handleTweetKeyEvent(event termbox.Event) {
	if handler.commandMode {
		return
	}

	switch GetActionByEvent(event) {
	case ACTION_RELOAD:
		handler.resetWithStatus("loading tweets...")
		handler.loadTweets()
	case ACTION_K, ACTION_UP:
		handler.container.UpSelectedTweet(handler.getTweets())
	case ACTION_J, ACTION_DOWN:
		handler.container.DownSelectedTweet(handler.getTweets())
	}

	handler.reset()
}

func (handler *Handler) handleCommandKeyEvent(event termbox.Event) {
	if !handler.commandMode {
		return
	}

	eventEmitter := views.GetCommandEventEmitter()

	switch GetActionByEvent(event) {
	case ACTION_LEFT:
		eventEmitter.Emit(views.CommandLeft)
	case ACTION_RIGHT:
		eventEmitter.Emit(views.CommandRight)
	case ACTION_DELETE:
		eventEmitter.Emit(views.CommandDelete)
	case ACTION_SPACE:
		eventEmitter.EmitWithValue(views.CommandAdd, []byte{' '})
	case ACTION_EXECUTE_COMMAND:
		err := handler.doCommand(handler.container.GetCommandValue())
		handler.commandMode = false
		eventEmitter.Emit(views.CommandEnd)
		if err != nil {
			handler.resetWithStatus(fmt.Sprintf("%s", err))
			return
		}
	default:
		if event.Ch != 0 {
			var u [utf8.UTFMax]byte
			s := utf8.EncodeRune(u[:], event.Ch)
			eventEmitter.EmitWithValue(views.CommandAdd, u[:s])
		}
	}

	handler.reset()
}

func (handler *Handler) draw() {
	handler.container.DrawTweets(handler.getTweets())
	handler.container.DrawStatus(handler.command.GetViewModeAsString(), "")
	handler.container.DrawCommand()
}

func (handler *Handler) drawWithStatus(status string) {
	handler.container.DrawTweets(handler.getTweets())
	handler.container.DrawStatus(handler.command.GetViewModeAsString(), status)
	handler.container.DrawCommand()
}

func (handler *Handler) reset() {
	handler.container.Clear()
	handler.drawWithStatus("@" + handler.consumer.GetScreenName())
	handler.container.Render()
}

func (handler *Handler) resetWithStatus(status string) {
	handler.container.Clear()
	handler.drawWithStatus(status)
	handler.container.Render()
}

func (handler *Handler) finish() {
	handler.container.Dispose()
}

func (handler *Handler) reload() {
	handler.resetWithStatus("loading tweets...")
	handler.loadTweets()
	handler.reset()
}

func (handler *Handler) getTweets() []*twitter.Tweet {
	viewMode := handler.command.GetViewMode()

	var tweets []*twitter.Tweet
	switch viewMode {
	case MODE_TIMELINE:
		tweets = GetTweetsStore().GetTimelineTweets()
	case MODE_MENTION:
		tweets = GetTweetsStore().GetMentionsTweets()
	case MODE_LIST:
		listName := handler.command.GetSlug()
		tweets = GetTweetsStore().GetListTweets(listName)
	}
	return tweets
}

func (handler *Handler) loadTweets() error {
	viewMode := handler.command.GetViewMode()

	switch viewMode {
	case MODE_TIMELINE:
		tweets, err := handler.consumer.GetTimeline()
		if err != nil {
			return err
		}
		GetTweetsStore().SetTimelineTweets(tweets)
	case MODE_MENTION:
		tweets, err := handler.consumer.GetMentions()
		if err != nil {
			return err
		}
		GetTweetsStore().SetMentionsTweets(tweets)
	case MODE_LIST:
		listName := handler.command.GetSlug()
		tweets, err := handler.consumer.GetListTimeline(listName)
		if err != nil {
			return err
		}
		GetTweetsStore().SetListTweets(listName, tweets)
	}

	return nil
}

func (handler *Handler) doCommand(value []byte) error {
	commandType, arg, err := handler.command.Parse(value)
	if err != nil {
		return err
	}

	switch commandType {
	case COMMAND_TIMELINE:
		handler.command.SetViewMode(MODE_TIMELINE)
	case COMMAND_MENTIONS:
		handler.command.SetViewMode(MODE_MENTION)
	case COMMAND_LIST:
		handler.command.SetViewMode(MODE_LIST)
		handler.command.SetSlug(arg)
	case COMMAND_TWEET:
		err = handler.consumer.Tweet(arg)
	case COMMAND_REPLY:
		tweet := handler.container.GetSelectedTweet(handler.getTweets())
		err = handler.consumer.Reply(arg, tweet)
	case COMMAND_FAVORITE:
		tweet := handler.container.GetSelectedTweet(handler.getTweets())
		err = handler.consumer.Favorite(tweet)
	case COMMAND_RETWEET:
		tweet := handler.container.GetSelectedTweet(handler.getTweets())
		err = handler.consumer.Retweet(tweet)
	}

	if err != nil {
		return err
	}

	handler.resetWithStatus("loading tweets...")
	handler.loadTweets()
	return nil
}
