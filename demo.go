package main

import (
	"fmt"
	"strconv"
)

func if_Leap_Year(year int) bool {
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

func if_30_Day_Month(month int) bool {
	return month == 4 || month == 6 || month == 9 || month == 11
}

func dateCheck(month int, day int, year int) bool {
	if year == 0 {
		return false
	}
	if !(month >= 1 && month <= 12) {
		return false
	} else if if_30_Day_Month(month) {
		return day >= 1 && day <= 30
	} else if month == 2 {
		if if_Leap_Year(year) {
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

	fmt.Println("Please enter you birthday(MM/DD/YYYY):")
	_, err := fmt.Scanf("%d/%d/%d", &month, &day, &year)

	if err == nil {
		if dateCheck(month, day, year) {
			fmt.Print("User input: ")
			fmt.Println(month, day, year)
		} else {
			fmt.Println("User input date no exist: " + strconv.Itoa(month) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(year))
		}
	} else {
		fmt.Println("User input error: ", err)
	}

	fmt.Print("Is it leap year: ")
	fmt.Println(if_Leap_Year(year))
}
