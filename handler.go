package main

import (
	"os"

	"github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
)

// Handler handle termbox events.
type Handler struct {
	args        []string
	apiClient   *twitter.Client
	quit        bool
	commandView *View
}

func NewHandler(args []string, apiClient *twitter.Client) *Looper {
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
				reset()
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
		xtermOffSequences = []string{
			// Ctrl + Arrow Keys
			"\x1b[1;5A", "\x1b[1;5B", "\x1b[1;5C", "\x1b[1;5D",
			// Shift + Arrow Left, Right
			"\x1b[1;2C", "\x1b[1;2D",
		}
		termbox.SetDisableEscSequence(xtermOffSequences)
	}

	termbox.Flush()

	return nil
}

func (self *Handler) setupView() {
	self.commandView = NewCommandView()
}

func (self *Handler) reset() {
	termbox.Clear(ColorBackground, ColorBackground)
	self.buffer.Draw()
	termbox.Flush()
}

func (self *Handler) finish() {
	self.apiClient.Close()
	termbox.Close()
}
