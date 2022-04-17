package calculate

import (
	"time"
)

func Calculation(startDay time.Time, endDay time.Time) (int, int, int) {
	year := endDay.Year()
	month := int(endDay.Month())
	day := endDay.Day()
	yearDiff := year - startDay.Year()

	monthDiff := month - int(startDay.Month())
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	dayDiff := day - startDay.Day()
	if dayDiff < 0 {
		dayDiff += 32 - dayDiff
		monthDiff--
	}

	return monthDiff, dayDiff, yearDiff

}
