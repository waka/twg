package main

import (
	"unicode/utf8"

	"github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
	"github.com/waka/twg/views"
)

// EventHandler
type Handler struct {
	args        []string
	apiClient   *twitter.Client
	container   *views.Container
	commandMode bool
	quit        bool
}

func NewHandler(args []string, apiClient *twitter.Client) *Handler {
	return &Handler{
		args:        args,
		apiClient:   apiClient,
		commandMode: false,
		quit:        false,
	}
}

func (handler *Handler) MainLoop() error {
	if err := handler.setupContainer(); err != nil {
		return err
	}
	defer handler.finish()

	handler.reset()
	if err := handler.loadTweet(); err != nil {
		return err
	}

	eventCh := make(chan termbox.Event)
	defer close(eventCh)

	go func() {
		for {
			eventCh <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case event := <-eventCh:
			if event.Type == termbox.EventResize {
				handler.reset()
			} else {
				if handler.commandMode && handler.getKeyEvent(event) != ACTION_QUIT {
					handler.handleCommandEvent(event)
				} else {
					handler.handleEvent(event)
				}
			}
		}
		if handler.quit {
			break
		}
	}

	return nil
}

func (handler *Handler) setupContainer() error {
	handler.container = views.NewContainer()
	if err := handler.container.Setup(); err != nil {
		return err
	}
	return nil
}

func (handler *Handler) reset() {
	handler.container.Render()
}

func (handler *Handler) finish() {
	//handler.apiClient.Close()
	handler.container.Dispose()
}

func (handler *Handler) handleEvent(event termbox.Event) {
	switch handler.getKeyEvent(event) {
	case ACTION_RELOAD:
		// refresh data and scroll top
	case ACTION_QUIT:
		// quit loop
		handler.quit = true
	case ACTION_UP:
		// select next tweet
	case ACTION_DOWN:
		// select prev tweet
	case ENTER_COMMAND_MODE:
		// delegate
		handler.handleCommandEvent(event)
	}
}

func (handler *Handler) handleCommandEvent(event termbox.Event) {
	eventEmitter := views.GetCommandEventEmitter()

	switch handler.getKeyEvent(event) {
	case ENTER_COMMAND_MODE:
		handler.commandMode = true
		eventEmitter.Emit(views.CommandStart)
	case ACTION_LEFT:
		eventEmitter.Emit(views.CommandLeft)
	case ACTION_RIGHT:
		eventEmitter.Emit(views.CommandRight)
	case ACTION_DELETE:
		eventEmitter.Emit(views.CommandDelete)
	case ACTION_EXECUTE_COMMAND:
		eventEmitter.Emit(views.CommandExecute)
	case ENTER_NORMAL_MODE:
		handler.commandMode = false
		eventEmitter.Emit(views.CommandEnd)
	default:
		if event.Ch != 0 {
			var u [utf8.UTFMax]byte
			s := utf8.EncodeRune(u[:], event.Ch)
			eventEmitter.EmitWithValue(views.CommandAdd, u[:s])
		}
	}

	handler.container.RenderCommand()
}

func (handler *Handler) getKeyEvent(event termbox.Event) Action {
	for _, keybind := range KeybindList {
		if event.Mod == keybind.Mod && event.Key == keybind.Key && event.Ch == keybind.Ch {
			return keybind.Action
		}
	}
	return NO_ACTION
}

func (handler *Handler) loadTweet() error {
	var (
		list []twitter.Tweet
		err  error
	)
	switch handler.container.GetViewMode() {
	case views.MODE_TIMELINE:
		list, err = handler.apiClient.GetTimeline()
	case views.MODE_LIST:
		list, err = handler.apiClient.GetTimeline()
	case views.MODE_MENTION:
		list, err = handler.apiClient.GetTimeline()
	}
	if err != nil {
		return err
	}
	handler.container.RenderContents()

	return nil
}
