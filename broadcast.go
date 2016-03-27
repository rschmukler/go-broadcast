package broadcast

import "sync"

// A Broadcaster is used to broadcast messages to multiple listeners at once
type Broadcaster struct {
	listeners []chan interface{}
	cap       int

	mu sync.Mutex
}

// NewBroadcaster builds a new generic Broadcaster. Cap is used as the buffer for each listener. If no cap
// is specified, a capacity of 0 will be used
func NewBroadcaster(capArg ...int) *Broadcaster {
	var cap int
	if len(capArg) == 0 {
		cap = 0
	} else {
		cap = capArg[0]
	}
	return &Broadcaster{
		listeners: []chan interface{}{},
		cap:       cap,
		mu:        sync.Mutex{},
	}
}

// Broadcast broadcasts a message to all listeners
func (b *Broadcaster) Broadcast(data interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()

	wg := sync.WaitGroup{}
	wg.Add(len(b.listeners))

	for _, listener := range b.listeners {
		// Send off to listeners in parallel
		go func(listener chan interface{}) {
			listener <- data
			wg.Done()
		}(listener)
	}

	wg.Wait()
}

// Listen registers a new listener for the broadcast
func (b *Broadcaster) Listen() <-chan interface{} {
	b.mu.Lock()
	defer b.mu.Unlock()

	listener := make(chan interface{}, b.cap)
	b.listeners = append(b.listeners, listener)

	return listener
}

// Close closes all listener channels
func (b *Broadcaster) Close() {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, listener := range b.listeners {
		close(listener)
	}
}

// Reset closes all listener channels and removes all known listeners from the broadcaster
func (b *Broadcaster) Reset() {
	b.mu.Lock()
	defer b.mu.Unlock()
	for _, listener := range b.listeners {
		close(listener)
	}
	b.listeners = []chan interface{}{}
}

// Remove removes a listener from the known listeners for a broadcaster, stopping it
// from receiving further broadcasts
func (b *Broadcaster) Remove(listener <-chan interface{}) {
	b.mu.Lock()
	defer b.mu.Unlock()
	for i, targetListener := range b.listeners {
		if targetListener == listener {
			b.listeners = append(b.listeners[:i], b.listeners[i+1:]...)
		}
	}
}
