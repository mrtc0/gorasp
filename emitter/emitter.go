package emitter

import (
	"sync"
)

type eventRegister struct {
	listeners map[string][]any
	mu        sync.RWMutex
}
type EventListener[O Operation, T any] func(O, T)

func addEventListener[O Operation, T any](r *eventRegister, name string, l EventListener[O, T]) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.listeners == nil {
		r.listeners = map[string][]any{}
	}

	r.listeners[name] = append(r.listeners[name], l)
}

func emitEvent[O Operation, T any](r *eventRegister, name string, op O, v T) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	if r.listeners == nil {
		return
	}

	for _, l := range r.listeners[name] {
		l.(EventListener[O, T])(op, v)
	}
}

func On[O Operation, E ArgOf[O]](name string, op Operation, l EventListener[O, E]) {
	addEventListener(&op.unwrap().eventRegister, name, l)
}
