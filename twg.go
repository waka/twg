package main

const (
	AppName = "twg"
	Version = "v0.2.0"
)

type ViewMode int

const (
	MODE_TIMELINE ViewMode = iota + 1
	MODE_MENTION
	MODE_LIST
)
