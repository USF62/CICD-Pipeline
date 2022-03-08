package main

import ("fmt"
		"strings"
		"time"
		"strconv"
)

func calculation (month, day, year, monthBirth, dayBirth, yearBirth int) (int, int, int){
	yearDiff := year - yearBirth

	monthDiff := month - monthBirth
	if monthDiff < 0 {
		monthDiff += 12
		yearDiff--
	}

	dayDiff := day - dayBirth
	if dayDiff < 0 {
		dayDiff += 32
		monthDiff--
	}

	return monthDiff, dayDiff, yearDiff

}

func handleInput(s, sep string) []string {
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


func main() {
	fmt.Println("enter date of birth (MM/DD/YYYY): ")
	todayDate := time.Now().Format("01/02/2006")
	var birthDate string
	fmt.Scanln(&birthDate)

	test := handleInput(todayDate, "/")
	month, _ := strconv.Atoi(test[0])
	day, _ := strconv.Atoi(test[1])
	year, _ := strconv.Atoi(test[2])

	test = handleInput(birthDate, "/")
	monthBirth, _ := strconv.Atoi(test[0])
	dayBirth, _ := strconv.Atoi(test[1])
	yearBirth, _ := strconv.Atoi(test[2])

	monthDiff, dayDiff, yearDiff := calculation(month, day, year, monthBirth, dayBirth, yearBirth)

	fmt.Printf("You are %d years, %d months, %d days old.", yearDiff, monthDiff, dayDiff)

}
