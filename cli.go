package main

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/jessevdk/go-flags"
)

type Options struct {
	Mode    string `short:"m" long:"mode" description:"you can select timeline(default) or mentions or list:slug"`
	Tweet   string `short:"t" long:"tweet" description:"update your status, and finish"`
	Reload  bool   `short:"r" long:"reload" description:"do auto reload by 2 minutes"`
	Version bool   `short:"v" long:"version"`
}

func (options *Options) GetMode() string {
	arr := strings.Split(options.Mode, ":")
	return arr[0]
}

func (options *Options) GetViewMode() ViewMode {
	mode := options.GetMode()
	switch mode {
	case "mentions":
		return MODE_MENTION
	case "list":
		return MODE_LIST
	default:
		return MODE_TIMELINE
	}
}

func (options *Options) GetSlug() string {
	arr := strings.Split(options.Mode, ":")
	if 2 > len(arr) {
		return ""
	}
	return arr[1]
}

// Exit codes are in value that represnet an exit code
// for a paticular error.
const (
	ExitCodeOK = 0 + iota

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeAuthError
	ExitCodeTweetError
	ExitCodeLoopError
)

func printErrorf(format string, args ...interface{}) {
	fmt.Println(fmt.Errorf(format, args...))
}

// CLI is empty struct.
type CLI struct{}

// Run is main function.
func (cli *CLI) Run(args []string) int {
	options, err := cli.parseArgs()
	if err != nil {
		return ExitCodeParseFlagsError
	}

	consumer := NewConsumer()
	if err := consumer.Auth(); err != nil {
		printErrorf("Failed to authenticate: %s", err)
		return ExitCodeAuthError
	}

	if options.Tweet != "" {
		if err := consumer.Tweet(options.Tweet); err != nil {
			printErrorf("Failed to tweet: %s", err)
			return ExitCodeTweetError
		}
		printErrorf("Done tweet: %s", options.Tweet)
		return ExitCodeOK
	} else {
		handler := NewHandler(NewCommand(options), consumer)
		if err := handler.MainLoop(); err != nil {
			printErrorf("Failed to loop: %s", err)
			return ExitCodeLoopError
		}
	}

	printErrorf("Bye")
	return ExitCodeOK
}

func (cli *CLI) parseArgs() (*Options, error) {
	options := &Options{}

	parser := flags.NewParser(options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		return nil, err
	}

	if options.Version {
		printErrorf("%s version: %s", AppName, Version)
		return nil, errors.New("Nothing to do")
	}

	if options.GetMode() == "list" && options.GetSlug() == "" {
		printErrorf("slug(list name) is required.\n")
		parser.WriteHelp(os.Stdout)
		return nil, errors.New("Nothing to do")
	}

	return options, nil
}
