package goatbot

import (
	"testing"
)

func TestAttendance(t *testing.T) {
	group := "Pirates"
	for _, p := range []string{"Tiago", "Sam", "Florian", "Jade", "Etienne", "Lesya", "Savy"} {
		Confirm(group, p)
	}
	for _, p := range []string{"Luc", "Tice"} {
		Cancel(group, p)
	}
	Cancel(group, "Jade")
	confirmed := ConfirmedList(group)
	if len(confirmed) != 6 {
		t.Error("Test failed")
	} else {
		t.Log(confirmed)
	}
}
