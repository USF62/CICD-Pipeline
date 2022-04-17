package util

import (
	"fmt"
	"time"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func FormatDate(year, month, day int) time.Time {
	loc, err := time.LoadLocation("America/Los_Angeles")
	CheckError(err)
	return time.Date(year, time.Month(month), day, 0, 0, 0, 0, loc)
}

func GetLocalTime() time.Time {
	loc, err := time.LoadLocation("America/Los_Angeles")
	CheckError(err)
	now := time.Now().UTC().In(loc)
	return now
}

func Is30DayMonth(month int) bool {
	return month == 4 || month == 6 || month == 9 || month == 11
}
