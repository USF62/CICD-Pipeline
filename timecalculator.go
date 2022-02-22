package main

import ("fmt"
		"strings"
		"time"
		"strconv"
)

func main() {
	fmt.Println("enter date of birth (MM/DD/YYYY): ")
	var birthDate string
	todayDate := time.Now().Format("01/02/2006")
	fmt.Scanln(&birthDate)

	test := strings.Split(todayDate, "/")
	month, _ := strconv.Atoi(test[0])
	day, _ := strconv.Atoi(test[1])
	year, _ := strconv.Atoi(test[2])

	test = strings.Split(birthDate, "/")
	monthBirth, _ := strconv.Atoi(test[0])
	dayBirth, _ := strconv.Atoi(test[1])
	yearBirth, _ := strconv.Atoi(test[2])


	yearDiff := year - yearBirth

	monthDiff := month - monthBirth
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	dayDiff := day - dayBirth
	if dayDiff < 0 {
		dayDiff = (32 - dayDiff)
		monthDiff--
	}

	fmt.Printf("You are %d years, %d months, %d days old.", yearDiff, monthDiff, dayDiff)

}
