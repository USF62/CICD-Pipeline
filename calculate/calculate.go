package calculate

import (
	"fmt"
	"strings"
)

func Calculation(month, day, year, monthBirth, dayBirth, yearBirth int) (int, int, int) {
	yearDiff := year - yearBirth
//change
	monthDiff := month - monthBirth
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

func HandleInput(s, sep string) []string {
	repeat := true
	var final []string

	for repeat == true {
		final = strings.Split(s, sep)
		if len(final) == 1 {
			fmt.Printf("Input incorrect! Please use %s as the seperators. \n", sep)
			fmt.Scanln(&s)
		} else {
			repeat = false
		}
	}
	return final
}
