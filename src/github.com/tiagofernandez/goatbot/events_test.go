package goatbot

import (
	"testing"
)

func TestAttendance(t *testing.T) {
	group := "Pirates"
	for _, p := range []string{"Tiago", "Sam", "Florian", "Jade", "Etienne", "Lesya", "Savy"} {
		Going(group, p)
	}
	for _, p := range []string{"Luc", "Tice"} {
		NotGoing(group, p)
	}
	NotGoing(group, "Jade")
	attending := AttendingList(group)
	if len(attending) != 6 {
		t.Error("Test failed")
	} else {
		t.Log(attending)
	}
}
