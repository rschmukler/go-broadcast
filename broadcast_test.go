package broadcast

import (
	"sync"
	"testing"
)

func TestBroadcaster(t *testing.T) {
	broadcaster := NewBroadcaster()

	var (
		recvCount     int
		listenerCount int
		data          = 1
	)

	wg := sync.WaitGroup{}

	addListener := func() {
		listenerCount++
		wg.Add(1)
		recv := <-broadcaster.Listen()
		if recv != data {
			t.Fatalf("Received incorrect data.\n\tExpected: %v\n\tReceived: %v", data, recv)
		}
		recvCount++
	}

	go addListener()
	go addListener()
	wg.Wait()

	if recvCount != listenerCount {
		t.Fatal("Received incorrect number of events")
	}

	broadcaster.Reset()
}

func TestBroadcaster_Reset(t *testing.T) {
	broadcaster := NewBroadcaster()

	broadcaster.Listen()
	broadcaster.Listen()
	broadcaster.Reset()

	// Since capacity is 0, this would block if there were listeners
	broadcaster.Broadcast(1)
}

func TestBroadcaster_Remove(t *testing.T) {
	broadcaster := NewBroadcaster()

	listener := broadcaster.Listen()

	go func() {
		<-listener
		t.Fatal("Listener received broadcast after being removed")
	}()

	broadcaster.Remove(listener)
	broadcaster.Broadcast(1)
}
