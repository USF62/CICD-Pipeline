package daysToBday

import (
	"fmt"
	"github.com/USF62/CICD-Pipeline/util"
	"math"
	"time"
)

/*
	Calculates the days to birthday from the current date and time
*/
func getDaysToBday(yearBD int, birthDate time.Time, now time.Time) int {
	nextBD := util.FormatDate(yearBD, int(birthDate.Month()), int(birthDate.Day()))
	daysToBD := math.Ceil(nextBD.Sub(now).Hours() / 24)
	fmt.Printf("You have around %.0f days until your birthday.\n", daysToBD)
	return int(daysToBD)
}

/*
	Finds the number of days to Birthday based on if birthday has passed in the current year (Used for Testing)
*/
func CalculateDaysToBday(birthDate time.Time) int {
	now := util.GetLocalTime()
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

/*
	Finds the number of days to Birthday based on if birthday has passed in the current year  (Used for Web Application)
*/
func CalculateDaysToBdayStr(birthDate time.Time) string {
	now := util.GetLocalTime()
	if (birthDate.Month() == now.Month() && birthDate.Day() >= now.Day()) || (birthDate.Month() > now.Month()) {
		if birthDate.Day() == now.Day() {
			return "Today is your birthday. Happy Birthday!"
		}
		return fmt.Sprintf("Your birthday this year is %d day(s) away", getDaysToBday(now.Year(), birthDate, now))
	} else if (birthDate.Month() == now.Month() && birthDate.Day() < now.Day()) || (birthDate.Month() < now.Month()) {
		fmt.Println("Your next birthday is next year!")
		yearBD := int(now.Year()) + 1
		return fmt.Sprintf("Your birthday next year is %d day(s) away!", getDaysToBday(yearBD, birthDate, now))
	}
	return ""
}
