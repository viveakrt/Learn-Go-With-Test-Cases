package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Vivek")
	want := "Hello, Vivek"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
