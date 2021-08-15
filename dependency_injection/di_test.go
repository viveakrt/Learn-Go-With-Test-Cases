package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Vivek")
	got := buffer.String()
	want := "Hello, Vivek"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
