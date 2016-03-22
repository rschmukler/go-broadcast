package broadcast

import (
	"errors"
	"testing"
)

func TestErrorBroadcaster(t *testing.T) {
	bc := NewErrorBroadcaster()

	var (
		inA             = errors.New("An error")
		inB, outA, outB error
	)

	done := make(chan bool)
	listener := bc.Listen()
	go func() {
		outA = <-listener
		outB = <-listener
		done <- true
	}()
	bc.Broadcast(inA)
	bc.Broadcast(inB)
	<-done

	if outA != inA {
		t.Fatal("Error was not passed through")
	}

	if outB != nil {
		t.Fatal("Nil error did not passthrough")
	}
}
