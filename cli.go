package main

import (
	"errors"
	"flag"
	"fmt"
)

// Exit codes are in value that represnet an exit code
// for a paticular error.
const (
	ExitCodeOK = 0 + iota

	// Errors start at 10
	ExitCodeError = 10 + iota
	ExitCodeParseFlagsError
	ExitCodeAuthError
	ExitCodeRenderingProcessError
)

func printErrorf(format string, args ...interface{}) {
	fmt.Println(fmt.Errorf(format, args...))
}

// CLI is empty struct.
type CLI struct {
}

// Run is main function.
func (cli *CLI) Run(args []string) int {
	args, err := cli.parseArgs(args)
	if err != nil {
		return ExitCodeParseFlagsError
	}

	accountManager := NewAccountManager()
	if err := accountManager.Auth(); err != nil {
		printErrorf("Failed to authenticate: %s", err)
		return ExitCodeAuthError
	}

	handler := NewHandler(args, accountManager.GetAPIClient())
	if err := handler.MainLoop(); err != nil {
		printErrorf("Failed to render: %s", err)
		return ExitCodeRenderingProcessError
	}

	printErrorf("Bye")
	return ExitCodeOK
}

func (cli *CLI) parseArgs(args []string) ([]string, error) {
	var version bool

	flags := flag.NewFlagSet(AppName, flag.ContinueOnError)
	flags.BoolVar(&version, "version", false, "")
	flags.BoolVar(&version, "v", false, "")

	// parse flag
	if err := flags.Parse(args[1:]); err != nil {
		printErrorf("Failed to parse args: %s", err)
		return nil, err
	}

	if version {
		printErrorf("%s version: %s", AppName, Version)
		return nil, errors.New("Nothing to do")
	}

	return flags.Args(), nil
}
