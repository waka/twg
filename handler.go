package main

import (
	"os"

	"github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
	"github.com/waka/twg/views"
)

// Handler handle termbox events.
type Handler struct {
	args        []string
	apiClient   *twitter.Client
	quit        bool
	contentView *views.CommandView
	commandView *views.CommandView
}

func NewHandler(args []string, apiClient *twitter.Client) *Handler {
	return &Handler{args: args, apiClient: apiClient, quit: false}
}

func (self *Handler) MainLoop() error {
	defer func() {
		self.finish()
	}()

	if err := self.setupTermbox(); err != nil {
		return err
	}
	if err := self.setupView(); err != nil {
		return err
	}

	ch := make(chan termbox.Event)
	go func() {
		for {
			ch <- termbox.PollEvent()
		}
	}()

	for {
		select {
		case event := <-ch:
			if event.Type == termbox.EventResize {
				self.reset()
			}
		}
		if self.quit {
			break
		}
	}

	return nil
}

func (self *Handler) setupTermbox() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputAlt)

	if os.Getenv("TERM") == "xterm" {
		/*
			xtermOffSequences := []string{
				// Ctrl + Arrow Keys
				"\x1b[1;5A", "\x1b[1;5B", "\x1b[1;5C", "\x1b[1;5D",
				// Shift + Arrow Left, Right
				"\x1b[1;2C", "\x1b[1;2D",
			}
			termbox.SetDisableEscSequence(xtermOffSequences)
		*/
	}

	termbox.Flush()

	return nil
}

func (self *Handler) setupView() error {
	self.commandView = views.NewCommandView()
	return nil
}

func (self *Handler) reset() {
	termbox.Clear(views.ColorBackground, views.ColorBackground)
	self.commandView.Draw()
	termbox.Flush()
}

func (self *Handler) finish() {
	//self.apiClient.Close()
	termbox.Close()
}
