package events

import (
	"errors"
	"sync"
)

type EventDispatcher struct {
	handlers map[string][]EventHandlerInterface
}

func NewEventDispatcher() *EventDispatcher {
	return &EventDispatcher{
		handlers: make(map[string][]EventHandlerInterface),
	}
}

func (ev *EventDispatcher) Dispatch(event EventInterface) error {
	if handlers, ok := ev.handlers[event.GetNome()]; ok {
		wg := &sync.WaitGroup{}
		for _, handler := range handlers {
			wg.Add(1)
			go handler.Handle(event, wg)
		}
		wg.Wait()
	}
	return nil
}

func (ev *EventDispatcher) RegisterHandler(eventName string, handler EventHandlerInterface) error {
	if _, ok := ev.handlers[eventName]; ok {
		for _, h := range ev.handlers[eventName] {
			if h == handler {
				return errors.New("handler already register")
			}
		}
	}

	ev.handlers[eventName] = append(ev.handlers[eventName], handler)
	return nil
}

func (ev *EventDispatcher) HasHandlers(eventName string, handler EventHandlerInterface) bool {
	if _, ok := ev.handlers[eventName]; ok {
		for _, h := range ev.handlers[eventName] {
			if h == handler {
				return true
			}
		}
	}

	return false
}

func (ev *EventDispatcher) Clear() error {
	ev.handlers = make(map[string][]EventHandlerInterface)

	return nil
}

func (ev *EventDispatcher) Remove(eventName string, handler EventHandlerInterface) error {
	if _, ok := ev.handlers[eventName]; ok {
		for i, h := range ev.handlers[eventName] {
			if h == handler {
				ev.handlers[eventName] = append(ev.handlers[eventName][:i], ev.handlers[eventName][i+1:]...)
			}
		}
	}

	return nil
}
