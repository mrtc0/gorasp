package emitter

import "sync"

var (
	eventEmitter *EventEmitter
	once         sync.Once
)

type EventEmitter struct {
	listeners map[string][]func(Params)
	mu        sync.RWMutex
}

func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]func(Params)),
	}
}

func (e *EventEmitter) On(event string, listener func(Params)) {
	e.mu.Lock()
	defer e.mu.Unlock()

	e.listeners[event] = append(e.listeners[event], listener)
}

func (e *EventEmitter) Emit(event string, params Params) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	for _, listener := range e.listeners[event] {
		listener(params)
	}
}

func GetEventEmitter() *EventEmitter {
	once.Do(func() {
		eventEmitter = NewEventEmitter()
	})

	return eventEmitter
}
