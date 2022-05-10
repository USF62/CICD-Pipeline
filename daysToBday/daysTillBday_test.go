package daysToBday

import (
	"fmt"
	"github.com/USF62/CICD-Pipeline/util"
	"testing"
	"time"
)

func getDateFromTime(days int) time.Time {
	now := util.GetLocalTime()
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
	result := CalculateDaysToBday(sameYearBday)
	assertEqual(t, result, 4, "The outputted result does not match the expected result.")
}

func TestUpcomingBirthdayNextYear(t *testing.T) {
	nextYearBday := getDateFromTime(-4)
	result := CalculateDaysToBday(nextYearBday)
	assertEqual(t, result, 361, "The outputted result does not match the expected result.")
}

func TestBirthdayToday(t *testing.T) {
	today := getDateFromTime(0)
	result := CalculateDaysToBday(today)
	assertEqual(t, result, 0, "The outputted result does not match the expected result.")
}
