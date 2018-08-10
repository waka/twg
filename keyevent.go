package main

import termbox "github.com/nsf/termbox-go"

type ActionType uint8
type Action uint8
type Keybind struct {
	Mod        termbox.Modifier
	Key        termbox.Key
	Ch         rune
	ActionType ActionType
	Action     Action
}

// ActionTypes
const (
	ACTION_TYPE_TWEET = iota + 1
	ACTION_TYPE_COMMAND
	ACTION_TYPE_BOTH
	ACTION_TYPE_ALL
)

// Actions
const (
	ACTION_RELOAD = iota + 1
	ACTION_QUIT
	ACTION_ESC
	ACTION_K
	ACTION_J
	ACTION_G
	ACTION_SHIFT_G
	ACTION_UP
	ACTION_DOWN
	ACTION_LEFT
	ACTION_RIGHT
	ACTION_DELETE
	ACTION_SPACE
	ACTION_EXECUTE_COMMAND
	ENTER_NORMAL_MODE
	ENTER_COMMAND_MODE
)

const NO_MOD = 0
const NO_KEY = 0
const NO_CH = 0
const NO_ACTION_TYPE = 0
const NO_ACTION = 0

var KeybindList = []Keybind{
	{NO_MOD, termbox.KeyCtrlQ, NO_CH, ACTION_TYPE_ALL, ACTION_QUIT},
	{NO_MOD, termbox.KeyCtrlC, NO_CH, ACTION_TYPE_ALL, ENTER_NORMAL_MODE},
	{NO_MOD, termbox.KeyEsc, NO_CH, ACTION_TYPE_ALL, ACTION_ESC},
	{NO_MOD, NO_KEY, ':', ACTION_TYPE_ALL, ENTER_COMMAND_MODE},
	{NO_MOD, termbox.KeyCtrlR, NO_CH, ACTION_TYPE_TWEET, ACTION_RELOAD},
	{NO_MOD, termbox.KeyArrowUp, NO_CH, ACTION_TYPE_TWEET, ACTION_UP},
	{NO_MOD, termbox.KeyArrowDown, NO_CH, ACTION_TYPE_TWEET, ACTION_DOWN},
	{NO_MOD, NO_KEY, 'k', ACTION_TYPE_BOTH, ACTION_K},
	{NO_MOD, NO_KEY, 'j', ACTION_TYPE_BOTH, ACTION_J},
	{NO_MOD, NO_KEY, 'g', ACTION_TYPE_BOTH, ACTION_G},
	{NO_MOD, NO_KEY, 'G', ACTION_TYPE_BOTH, ACTION_SHIFT_G},
	{NO_MOD, termbox.KeyArrowLeft, NO_CH, ACTION_TYPE_COMMAND, ACTION_LEFT},
	{NO_MOD, termbox.KeyArrowRight, NO_CH, ACTION_TYPE_COMMAND, ACTION_RIGHT},
	{NO_MOD, termbox.KeyBackspace, NO_CH, ACTION_TYPE_COMMAND, ACTION_DELETE},
	{NO_MOD, termbox.KeyBackspace2, NO_CH, ACTION_TYPE_COMMAND, ACTION_DELETE},
	{NO_MOD, termbox.KeySpace, NO_CH, ACTION_TYPE_COMMAND, ACTION_SPACE},
	{NO_MOD, termbox.KeyEnter, NO_CH, ACTION_TYPE_COMMAND, ACTION_EXECUTE_COMMAND},
}

func GetActionTypeByEvent(event termbox.Event) ActionType {
	for _, keybind := range KeybindList {
		if event.Mod == keybind.Mod && event.Key == keybind.Key && event.Ch == keybind.Ch {
			return keybind.ActionType
		}
	}
	return NO_ACTION_TYPE
}

func GetActionByEvent(event termbox.Event) Action {
	for _, keybind := range KeybindList {
		if event.Mod == keybind.Mod && event.Key == keybind.Key && event.Ch == keybind.Ch {
			return keybind.Action
		}
	}
	return NO_ACTION
}
