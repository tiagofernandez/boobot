package boobot

import "testing"

func TestGimmeSomeBoobs(t *testing.T) {
	fail := "http://tinyurl.com/tits-not-found"
	actual := GimmeSomeBoobs()
	if actual == fail {
		t.Error("Test failed: " + actual)
	} else {
		t.Log(actual)
	}
}
