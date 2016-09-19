package goatbot

import (
	"testing"
)

func TestAttendance(t *testing.T) {
	group := "Pirates"
	Going(group, "Tiago")
	Going(group, "Sam")
	NotGoing(group, "Luc")
	NotGoing(group, "Tice")
	Going(group, "Florian")
	Going(group, "Jade")
	Going(group, "Etienne")
	Going(group, "Lesya")
	Going(group, "Savvy")
	NotGoing(group, "Jade")
	attending := AttendingList(group)
	if len(attending) != 6 {
		t.Error("Test failed")
	} else {
		t.Log(attending)
	}
}
