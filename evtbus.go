package evtbus

import "sync"

//Event stores event topic and data
type Event struct {
	Topic string
	Data  interface{}
}

//EventChannel aliases a chan Event
type EventChannel chan Event

//EventBus stores the subscribers of a topic
type EventBus struct {
	subscribers map[string][]EventChannel
	rm          sync.RWMutex
}

//New creates an EventBus with empty map of subscribers
func New() *EventBus {
	return &EventBus{
		subscribers: map[string][]EventChannel{},
	}
}

//Subscribe an EventChannel to a topic. When an Event is published all channels will receive it.
func (eb *EventBus) Subscribe(topic string, ch EventChannel) {
	eb.rm.Lock()
	if subs, found := eb.subscribers[topic]; found {
		eb.subscribers[topic] = append(subs, ch)
	} else {
		eb.subscribers[topic] = append([]EventChannel{}, ch)
	}
	eb.rm.Unlock()
}

//Publish an Event to all subscribers in a topic
func (eb *EventBus) Publish(event Event) {
	eb.rm.RLock()
	if chans, found := eb.subscribers[event.Topic]; found {
		subs := append([]EventChannel{}, chans...)
		go func(event Event, subs []EventChannel) {
			for _, ch := range subs {
				ch <- event
			}
		}(event, subs)
	}
	eb.rm.RUnlock()
}
