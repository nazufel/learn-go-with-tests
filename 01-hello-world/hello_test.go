package main

import (
	"fmt"
	"testing"
)

func TestHello(t *testing.T) {
	got := Hello("Ryan")
	want := "Hello, Ryan"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	fmt.Printf("got %q want %q", got, want)
}
