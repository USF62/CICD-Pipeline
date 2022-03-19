package DaysToBday

import (
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func formatDate(year, month, date int) time.Time {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	return time.Date(year, time.Month(month), date, 0, 0, 0, 0, loc)
}

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}
func splitBday(bDay string) time.Time {
	splitBirthDay := strings.Split(bDay, "/")
	month, error := strconv.Atoi(splitBirthDay[0])
	CheckError(error)
	day, error := strconv.Atoi(splitBirthDay[1])
	CheckError(error)
	year, error := strconv.Atoi(splitBirthDay[2])
	CheckError(error)

	birthDate := formatDate(year, month, day)
	return birthDate
}

func calculateDaysToBday(birthDate time.Time) int {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc)

	if (birthDate.Month() == now.Month() && birthDate.Day() >= now.Day()) || (birthDate.Month() > now.Month()) {
		if birthDate.Day() == now.Day() {
			fmt.Println("Today is your birthday. Happy Birthday!")
			return 0
		}
		fmt.Println("Your next birthday is this year!")
		nextBD := formatDate(int(now.Year()), int(birthDate.Month()), int(birthDate.Day()))
		daysToBD := math.Ceil(nextBD.Sub(now).Hours() / 24)
		fmt.Printf("You have around %.0f days until your birthday.\n", daysToBD)
		return int(daysToBD)

	} else if (birthDate.Month() == now.Month() && birthDate.Day() < now.Day()) || (birthDate.Month() < now.Month()) {
		fmt.Println("Your next birthday is next year!")
		yearBD := int(now.Year()) + 1
		nextBD := formatDate(yearBD, int(birthDate.Month()), int(birthDate.Day()))
		daysToBD := math.Ceil(nextBD.Sub(now).Hours() / 24)
		fmt.Printf("You have around %.0f days until your birthday.\n", daysToBD)
		return int(daysToBD)
	}
	return 0
}
