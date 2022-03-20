package DaysToBday

import (
	"fmt"
	"math"
	"time"
)

func getDaysToBday(yearBD int, birthDate time.Time, now time.Time) int {
	nextBD := formatDate(yearBD, int(birthDate.Month()), int(birthDate.Day()))
	daysToBD := math.Ceil(nextBD.Sub(now).Hours() / 24)
	fmt.Printf("You have around %.0f days until your birthday.\n", daysToBD)
	return int(daysToBD)
}

func calculateDaysToBday(birthDate time.Time) int {
	now := getLocalTime()
	if (birthDate.Month() == now.Month() && birthDate.Day() >= now.Day()) || (birthDate.Month() > now.Month()) {
		if birthDate.Day() == now.Day() {
			fmt.Println("Today is your birthday. Happy Birthday!")
			return 0
		}
		fmt.Println("Your next birthday is this year!")
		return getDaysToBday(now.Year(), birthDate, now)
	} else if (birthDate.Month() == now.Month() && birthDate.Day() < now.Day()) || (birthDate.Month() < now.Month()) {
		fmt.Println("Your next birthday is next year!")
		yearBD := int(now.Year()) + 1
		return getDaysToBday(yearBD, birthDate, now)
	}
	return 0
}
