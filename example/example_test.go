package example

import "testing"

func TestGreeting(t *testing.T) {
	expected := "Hey you, let's Go!"
	msg := Greeting("you")

	if expected != msg {
		t.Errorf("Expected: %s; got: %s", expected, msg)
	}
}
