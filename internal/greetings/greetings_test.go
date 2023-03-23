package greetings_test

import (
	"testing"

	"github.com/sergiofgonzalez/greeter/internal/greetings"
)

func TestHello(t *testing.T) {
	want := "Hello, World!"
	got := greetings.Hello()

	if got != want {
		t.Errorf("got %q, but wanted %q", got, want)
	}
}