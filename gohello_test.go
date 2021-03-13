package gohello

import (
	"testing"
)

func TestMessage(t *testing.T) {
	if msg := Message(); msg != "Hello!" {
		t.Error("Message() =", msg)
	}
}
