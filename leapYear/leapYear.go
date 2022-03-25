package main

import (
	"fmt"
	"strconv"
)

func IsLeapYear(year int) bool {
	/* Using the new leap year rule since 1582*/
	return ((year%4 == 0 && year%100 != 0) || year%400 == 0) && year >= 1582
}

func Is30DayMonth(month int) bool {
	return month == 4 || month == 6 || month == 9 || month == 11
}

func IsValidDate(month int, day int, year int) bool {
	/* Using the new leap year rule since 1582*/
	if year < 1582 {
		return false
	}
	if year <= 0 || !(month >= 1 && month <= 12) || day < 1 {
		return false
	} else if Is30DayMonth(month) {
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

func main() {

	var month, day, year int

	fmt.Println("Please enter your birthday(MM/DD/YYYY):")
	_, err := fmt.Scanf("%d/%d/%d", &month, &day, &year)

	if err == nil {
		if IsValidDate(month, day, year) {
			fmt.Print("User input: ")
			fmt.Println(month, day, year)
		} else {
			fmt.Println("User input date no exist: " + strconv.Itoa(month) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(year))
		}
	} else {
		fmt.Println("User input error: ", err)
	}

	fmt.Print("Is it leap year: ")
	fmt.Println(IsLeapYear(year))
}
