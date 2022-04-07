package calculate

import (
	"github.com/venkatrajm815/USF62-Pipeline/util"
	"time"
)

func Calculation(monthBirth, dayBirth, yearBirth int) (int, int, int) {
	currentdate := util.GetLocalTime()
	var month time.Month
	var day, year int
	year, month, day = currentdate.Date()
	yearDiff := year - yearBirth

	monthDiff := int(month) - monthBirth
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	dayDiff := day - dayBirth
	if dayDiff < 0 {
		dayDiff += 32 - dayDiff
		monthDiff--
	}

	return monthDiff, dayDiff, yearDiff

}
