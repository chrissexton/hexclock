package hextime

import (
	"testing"
	"time"
)

func check(t *testing.T, actual, expected string) {
	if actual != expected {
		t.Errorf("%s != %s", actual, expected)
	}
}

func getTime(when string) time.Time {
	t, _ := time.ParseInLocation(time.Kitchen, when, time.Local)
	return t
}

func TestCalcTime(t *testing.T) {
	check(t, Convert(getTime("7:30AM")), ".5000")
	check(t, Convert(getTime("9:00AM")), ".6000")
	check(t, Convert(getTime("12:00PM")), ".8000")
	check(t, Convert(getTime("5:00PM")), ".B555")
	check(t, Convert(getTime("9:00PM")), ".E000")
}
