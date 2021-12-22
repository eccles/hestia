package main

import (
	"testing"
)

func TestHello(t *testing.T) {
	t.Parallel()

	want := "Don't communicate by sharing memory, share memory by communicating."

	if got := Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
