package boobot

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := "👋, 🌎!"
	actual := HelloWorld()
	if actual != expected {
		t.Error("Test failed")
	}
}
