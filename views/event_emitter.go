package views

import (
	"reflect"
	"sync"
)

type EventEmitter struct {
	eventMutex     *sync.Mutex
	eventListeners []func(event CommandEvent)
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{eventMutex: new(sync.Mutex)}
}

func (e *EventEmitter) EmitEvent(event CommandEvent) {
	e.eventMutex.Lock()
	listeners := make([]func(event CommandEvent), len(e.eventListeners))
	copy(listeners, e.eventListeners)
	e.eventMutex.Unlock()

	for _, l := range listeners {
		l(event)
	}
}

func (e *EventEmitter) AddEventListener(listener func(event CommandEvent)) {
	e.eventMutex.Lock()
	e.eventListeners = append(e.eventListeners, listener)
	e.eventMutex.Unlock()
}

func (e *EventEmitter) RemoveEventListener(listener func(event CommandEvent)) {
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
