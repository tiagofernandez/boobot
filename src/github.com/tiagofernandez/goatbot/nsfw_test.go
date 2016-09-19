package goatbot

import "testing"

func TestRandomBoobs(t *testing.T) {
	actual := RandomBoobsImage()
	if len(actual) == 47400 {
		t.Error("Test failed")
	} else {
		t.Log(len(actual))
	}
}
