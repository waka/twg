package main

import (
	"errors"
	"regexp"
)

type CommandType uint8

const (
	COMMAND_TIMELINE = iota + 1
	COMMAND_MENTIONS
	COMMAND_LIST
	COMMAND_TWEET
	COMMAND_REPLY
	COMMAND_FAVORITE
	COMMAND_RETWEET
	COMMAND_NOT_FOUND
)

var re = regexp.MustCompile(`:([a-z]*)\s?(.*)`)

type Command struct {
	viewMode ViewMode
	slug     string
	reload   bool
}

func NewCommand(options *Options) *Command {
	return &Command{
		viewMode: options.GetViewMode(),
		slug:     options.GetSlug(),
		reload:   options.Reload,
	}
}

func (command *Command) GetViewMode() ViewMode {
	return command.viewMode
}

func (command *Command) SetViewMode(viewMode ViewMode) {
	command.viewMode = viewMode
}

func (command *Command) GetViewModeAsString() string {
	switch command.viewMode {
	case MODE_TIMELINE:
		return "Timeline"
	case MODE_MENTION:
		return "Mentions"
	case MODE_LIST:
		return "List: " + command.slug
	}
	return ""
}

func (command *Command) GetSlug() string {
	return command.slug
}

func (command *Command) SetSlug(slug string) {
	command.slug = slug
}

func (command *Command) IsReloadable() bool {
	return command.reload
}

func (command *Command) Parse(value []byte) (CommandType, string, error) {
	var (
		commandType CommandType
		err         error
	)
	result := re.FindAllStringSubmatch(string(value), -1)
	action := result[0][1]
	arg := result[0][2]

	switch action {
	case "tl":
		commandType = COMMAND_TIMELINE
		if arg != "" {
			err = errors.New("Argument is not required")
		}
	case "mentions":
		commandType = COMMAND_MENTIONS
		if arg != "" {
			err = errors.New("Argument is not required")
		}
	case "list":
		commandType = COMMAND_LIST
		if arg == "" {
			err = errors.New("List name is required")
		}
	case "tweet":
		commandType = COMMAND_TWEET
		if arg == "" {
			err = errors.New("text is required")
		}
	case "reply":
		commandType = COMMAND_REPLY
		if arg == "" {
			err = errors.New("text is required")
		}
	case "fav":
		commandType = COMMAND_FAVORITE
	case "rt":
		commandType = COMMAND_RETWEET
	default:
		commandType = COMMAND_NOT_FOUND
		err = errors.New("Command is not found")
	}

	return commandType, arg, err
}
