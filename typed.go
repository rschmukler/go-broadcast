package broadcast

// Typed broadcasters for convenience

// ErrorBroadcaster is a broadcaster which broadcasts errors.
type ErrorBroadcaster struct {
	*Broadcaster
}

// NewErrorBroadcaster builds a broadcaster for broadcasting errors
func NewErrorBroadcaster(cap ...int) *ErrorBroadcaster {
	return &ErrorBroadcaster{
		Broadcaster: NewBroadcaster(cap...),
	}
}

// Listen registers a new listener for the broadcast
func (b *ErrorBroadcaster) Listen() <-chan error {
	rawListener := b.Broadcaster.Listen()
	listener := make(chan error, b.Broadcaster.cap)
	go func() {
		for data := range rawListener {
			if data != nil {
				listener <- data.(error)
			} else {
				listener <- nil
			}
		}
		close(listener)
	}()
	return listener
}

// Broadcast broadcasts an error to all listeners
func (b *ErrorBroadcaster) Broadcast(err error) {
	b.Broadcaster.Broadcast(err)
}

// BoolBroadcaster is a broadcaster which broadcasts booleans.
type BoolBroadcaster struct {
	*Broadcaster
}

// NewBoolBroadcaster builds a broadcaster for broadcasting booleans
func NewBoolBroadcaster(cap ...int) *BoolBroadcaster {
	return &BoolBroadcaster{
		Broadcaster: NewBroadcaster(cap...),
	}
}

// Listen registers a new listener for the broadcast
func (b *BoolBroadcaster) Listen() <-chan bool {
	rawListener := b.Broadcaster.Listen()
	listener := make(chan error, b.Broadcaster.cap)
	go func() {
		for data := range rawListener {
			listener <- data.(bool)
		}
		close(listener)
	}()
	return listener
}

// Broadcast broadcasts an bool to all listeners
func (b *BoolBroadcaster) Broadcast(val bool) {
	b.Broadcaster.Broadcast(val)
}
