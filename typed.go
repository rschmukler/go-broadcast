package broadcast

// Typed broadcasters for convenience

// ErrorBroadcaster is a broadcaster which broadcasts errors
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
