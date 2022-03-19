package DaysToBday

import (
	"fmt"
	"testing"
	"time"
)

func getDateFromTime(days int) time.Time {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc)
	date := now.AddDate(0, 0, days)
	return date
}

func assertEqual(t *testing.T, a interface{}, b interface{}, message string) {
	if a == b {
		return
	}
	if len(message) == 0 {
		message = fmt.Sprintf("%v != %v", a, b)
	}
	t.Fatal(message)
}
func TestUpcomingBirthdaySameYear(t *testing.T) {
	sameYearBday := getDateFromTime(4)
	result := calculateDaysToBday(sameYearBday)
	assertEqual(t, result, 4, "The outputted result does not match the expected result.")
}

func TestUpcomingBirthdayNextYear(t *testing.T) {
	nextYearBday := getDateFromTime(-4)
	result := calculateDaysToBday(nextYearBday)
	assertEqual(t, result, 361, "The outputted result does not match the expected result.")
}

func TestBirthdayToday(t *testing.T) {
	today := getDateFromTime(0)
	result := calculateDaysToBday(today)
	assertEqual(t, result, 0, "The outputted result does not match the expected result.")
}
