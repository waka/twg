package views

import (
	"reflect"
	"sync"
)

type CommandEventEmitter struct {
	eventMutex     *sync.Mutex
	eventListeners []func(event *CommandEvent)
}

var sharedInstance *CommandEventEmitter = newCommandEventEmitter()

func newCommandEventEmitter() *CommandEventEmitter {
	return &CommandEventEmitter{eventMutex: new(sync.Mutex)}
}

func GetCommandEventEmitter() *CommandEventEmitter {
	return sharedInstance
}

func (e *CommandEventEmitter) Emit(eventType CommandEventType) {
	event := NewCommandEvent(eventType, []byte{})
	e.emitEvent(event)
}

func (e *CommandEventEmitter) EmitWithValue(eventType CommandEventType, value []byte) {
	event := NewCommandEvent(eventType, value)
	e.emitEvent(event)
}

func (e *CommandEventEmitter) emitEvent(event *CommandEvent) {
	e.eventMutex.Lock()
	listeners := make([]func(event *CommandEvent), len(e.eventListeners))
	copy(listeners, e.eventListeners)
	e.eventMutex.Unlock()

	for _, l := range listeners {
		l(event)
	}
}

func (e *CommandEventEmitter) AddEventListener(listener func(event *CommandEvent)) {
	e.eventMutex.Lock()
	e.eventListeners = append(e.eventListeners, listener)
	e.eventMutex.Unlock()
}

func (e *CommandEventEmitter) RemoveEventListener(listener func(event *CommandEvent)) {
	listenerPtr := reflect.ValueOf(listener).Pointer()
	e.eventMutex.Lock()
	listeners := e.eventListeners[:0]
	for _, l := range e.eventListeners {
		if reflect.ValueOf(l).Pointer() != listenerPtr {
			listeners = append(listeners, l)
		}
	}
	e.eventListeners = listeners
	e.eventMutex.Unlock()
}
