package main

import (
	"fmt"
	"testing"
	"time"
)

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
func testUpcomingBirthdaySameYear(t *testing.T) {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc)
	future := now.AddDate(0, 0, 4)
	result := calculateDaysToBday(future)
	assertEqual(t, result, 4, "")
}

func testUpcomingBirthdayNextYear(t *testing.T) {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc)
	future := now.AddDate(0, 0, -4)
	result := calculateDaysToBday(future)
	assertEqual(t, result, -4, "")
}
