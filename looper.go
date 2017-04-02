package main

import (
	"github.com/nsf/termbox-go"
	"github.com/waka/twg/twitter"
)

type Looper struct {
	args      []string
	apiClient *twitter.Client
}

func NewLooper(args []string, apiClient *twitter.Client) *Looper {
	return &Looper{args: args, apiClient: apiClient}
}

func (self *Looper) MainLoop() error {
	defer func() {
		self.finish()
	}()

	if err := self.setupTermbox(); err != nil {
		return err
	}

	return nil
}

func (self *Looper) setupTermbox() error {
	if err := termbox.Init(); err != nil {
		return err
	}

	termbox.SetOutputMode(termbox.Output256)
	termbox.SetInputMode(termbox.InputAlt)
	//drawText("Now Loading...", 0, 0, ColorWhite, ColorBackground)
	termbox.Flush()

	return nil
}

func (self *Looper) finish() {
	termbox.Close()
}
