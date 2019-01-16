package main

// Subscription to observable
type Subscription struct {
	unsubscribe func()
}

// Unsubscribe to observble
func (s *Subscription) Unsubscribe() {
	if s.unsubscribe != nil {
		s.unsubscribe()
	}
}

// Observer manages values in an observable
type Observer struct {
}

// func (o *Observer) err() {
// 	o.err()
// }
// func (o *Observer) complete() {
// 	o.complete()
// }

// NewObservable creates new observable
func NewObservable(next func(o *Observable)) *Observable {
	observer := &Observer{}
	observable := &Observable{obs: observer}
	observable.valueProducer = next
	return observable
}

// Observable lets you watch thing for events
type Observable struct {
	obs           *Observer
	valueProducer func(o *Observable)
	subscribers   []func(int)
	isActive      bool
}

// Subscribe to events from observable, rest will be error and complete fn
func (o *Observable) Subscribe(fn func(int), rest ...func()) *Subscription {
	len := len(o.subscribers)
	o.subscribers = append(o.subscribers, fn)
	unsub := &Subscription{}
	unsub.unsubscribe = func() {
		o.subscribers[len] = nil
	}
	if !o.isActive {
		o.isActive = true
		o.valueProducer(o)
	}
	return unsub
}

// Next value sent to subscribers
func (o *Observable) Next(val int) {
	for _, subFn := range o.subscribers {
		subFn(val)
	}
}
