package goatbot

import "testing"

func TestGimmeSomeBoobs(t *testing.T) {
	actual := GimmeSomeBoobs()
	if len(actual) == 47400 {
		t.Error("Test failed")
	} else {
		t.Log(len(actual))
	}
}
