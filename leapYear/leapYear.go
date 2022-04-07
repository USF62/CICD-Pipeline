package leapYear

import (
	"github.com/venkatrajm815/USF62-Pipeline/util"
)

func IsLeapYear(year int) bool {
	/* Using the new leap year rule since 1582*/
	return ((year%4 == 0 && year%100 != 0) || year%400 == 0) && year >= 1582
}

func IsValidDate(month int, day int, year int) bool {
	/* Using the new leap year rule since 1582*/
	if year < 1582 {
		return false
	}
	if year <= 0 || !(month >= 1 && month <= 12) || day < 1 {
		return false
	} else if util.Is30DayMonth(month) {
		return day >= 1 && day <= 30
	} else if month == 2 {
		if IsLeapYear(year) {
			return day >= 1 && day <= 29
		} else {
			return day >= 1 && day <= 28
		}
	} else {
		return day >= 1 && day <= 31
	}
}
