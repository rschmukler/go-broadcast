package broadcast

// Typed broadcasters for convenience

// ErrorBroadcaster is a broadcaster which broadcasts errors.
type ErrorBroadcaster struct {
	*Broadcaster
	listeners []typedListener
}

type typedListener struct {
	raw      <-chan interface{}
	boolOut  <-chan bool
	errorOut <-chan error
	intOut   <-chan int
}

// NewErrorBroadcaster builds a broadcaster for broadcasting errors
func NewErrorBroadcaster(cap ...int) *ErrorBroadcaster {
	return &ErrorBroadcaster{
		Broadcaster: NewBroadcaster(cap...),
		listeners:   []typedListener{},
	}
}

// Listen registers a new listener for the broadcast
func (b *ErrorBroadcaster) Listen() <-chan error {
	rawListener := b.Broadcaster.Listen()
	out := make(chan error, b.Broadcaster.cap)
	b.listeners = append(b.listeners, typedListener{raw: rawListener, errorOut: out})
	go func() {
		for data := range rawListener {
			if data != nil {
				out <- data.(error)
			} else {
				out <- nil
			}
		}
		close(out)
	}()
	return out
}

// Remove removes a listener from the known listeners for a broadcaster, stopping it
// from receiving further broadcasts
func (b *ErrorBroadcaster) Remove(listener <-chan error) {
	for _, typedListener := range b.listeners {
		if listener == typedListener.errorOut {
			b.Broadcaster.Remove(typedListener.raw)
		}
	}
}

// Broadcast broadcasts an error to all listeners
func (b *ErrorBroadcaster) Broadcast(err error) {
	b.Broadcaster.Broadcast(err)
}

// BoolBroadcaster is a broadcaster which broadcasts booleans.
type BoolBroadcaster struct {
	*Broadcaster
	listeners []typedListener
}

// NewBoolBroadcaster builds a broadcaster for broadcasting booleans
func NewBoolBroadcaster(cap ...int) *BoolBroadcaster {
	return &BoolBroadcaster{
		Broadcaster: NewBroadcaster(cap...),
		listeners:   []typedListener{},
	}
}

// Listen registers a new listener for the broadcast
func (b *BoolBroadcaster) Listen() <-chan bool {
	rawListener := b.Broadcaster.Listen()
	out := make(chan bool, b.Broadcaster.cap)
	b.listeners = append(b.listeners, typedListener{raw: rawListener, boolOut: out})
	go func() {
		for data := range rawListener {
			out <- data.(bool)
		}
		close(out)
	}()
	return out
}

// Broadcast broadcasts an bool to all listeners
func (b *BoolBroadcaster) Broadcast(val bool) {
	b.Broadcaster.Broadcast(val)
}

// Remove removes a listener from the known listeners for a broadcaster, stopping it
// from receiving further broadcasts
func (b *BoolBroadcaster) Remove(listener <-chan bool) {
	for _, typedListener := range b.listeners {
		if listener == typedListener.boolOut {
			b.Broadcaster.Remove(typedListener.raw)
		}
	}
}

// IntBroadcaster is a broadcaster which broadcasts ints.
type IntBroadcaster struct {
	*Broadcaster
	listeners []typedListener
}

// NewIntBroadcaster builds a broadcaster for broadcasting ints
func NewIntBroadcaster(cap ...int) *IntBroadcaster {
	return &IntBroadcaster{
		Broadcaster: NewBroadcaster(cap...),
		listeners:   []typedListener{},
	}
}

// Listen registers a new listener for the broadcast
func (b *IntBroadcaster) Listen() <-chan int {
	rawListener := b.Broadcaster.Listen()
	out := make(chan int, b.Broadcaster.cap)
	b.listeners = append(b.listeners, typedListener{raw: rawListener, intOut: out})
	go func() {
		for data := range rawListener {
			out <- data.(int)
		}
		close(out)
	}()
	return out
}

// Broadcast broadcasts an int to all listeners
func (b *IntBroadcaster) Broadcast(val int) {
	b.Broadcaster.Broadcast(val)
}

// Remove removes a listener from the known listeners for a broadcaster, stopping it
// from receiving further broadcasts
func (b *IntBroadcaster) Remove(listener <-chan int) {
	for _, typedListener := range b.listeners {
		if listener == typedListener.intOut {
			b.Broadcaster.Remove(typedListener.raw)
		}
	}
}
