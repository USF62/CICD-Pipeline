package main

import (
	"fmt"
	calculate "github.com/venkatrajm815/USF62-Pipeline/ageCalculator"
	"github.com/venkatrajm815/USF62-Pipeline/daysToBday"
	"github.com/venkatrajm815/USF62-Pipeline/horoscope"
	"github.com/venkatrajm815/USF62-Pipeline/leapYear"
	_ "github.com/venkatrajm815/USF62-Pipeline/leapYear"
	"github.com/venkatrajm815/USF62-Pipeline/util"
	"strconv"
	"time"
)

func ScanInputDate() (int, int, int) {
	var month, day, year int
	fmt.Println("Please enter your birthday(MM/DD/YYYY):")
	_, err := fmt.Scanf("%d/%d/%d", &month, &day, &year)
	_, _ = fmt.Scanln()

	if err == nil {
		if leapYear.IsValidDate(month, day, year) {
			return month, day, year
		} else {
			fmt.Println("User input date no exist: " + strconv.Itoa(month) + "/" + strconv.Itoa(day) + "/" + strconv.Itoa(year))
			return 0, 0, 0
		}
	} else {
		fmt.Println("User input error: ", err)
		return 0, 0, 0
	}
}

func main() {
	var index string
	var month, day, year int
	run := true
	for run {
		menu := "\n\nWelcome to Age Calculator!\n" +
			"Input the index of the feature you want to use: (EX: 1)\n" +
			"1. Calculate your age\n" +
			"2. Calculate many days till your birthday\n" +
			"3. Is leap year\n" +
			"4. Your horoscope\n" +
			"5. Summarize\n" +
			"6. Exit"

		fmt.Println(menu)

		_, err := fmt.Scanln(&index)
		if err == nil {
			switch index {
			case "1":
				var rm, rd, ry int
				month, day, year = ScanInputDate()
				rm, rd, ry = calculate.Calculation(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
					util.GetLocalTime())
				fmt.Println("You are " + strconv.Itoa(ry) + " years " + strconv.Itoa(rm) + " months " + strconv.Itoa(rd) + " days old now.")
			case "2":
				month, day, year = ScanInputDate()
				daysToBday.CalculateDaysToBday(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
			case "3":
				var year int
				fmt.Println("Please enter a year: (EX: YYYY)")
				_, err := fmt.Scanf("%d", &year)
				if err != nil {
					fmt.Println("User enter non-integer.")
				}
				_, _ = fmt.Scanln()
				if leapYear.IsLeapYear(year) {
					fmt.Println(strconv.Itoa(year) + " is leap year.")
				} else {
					fmt.Println(strconv.Itoa(year) + " is not leap year.")
				}
			case "4":
				var horo, zodiac string
				month, day, year = ScanInputDate()
				horo = horoscope.GetHoroscope(month, day)
				zodiac = horoscope.GetZodiac(month, day, year)
				fmt.Println("Your horoscope is " + horo + " and your zodiac is " + zodiac)
			case "5":
				var rm, rd, ry int
				month, day, year = ScanInputDate()
				rm, rd, ry = calculate.Calculation(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC),
					util.GetLocalTime())
				fmt.Println("You are " + strconv.Itoa(ry) + " years " + strconv.Itoa(rm) + " months " + strconv.Itoa(rd) + " days old now.")
				daysToBday.CalculateDaysToBday(time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC))
				if leapYear.IsLeapYear(year) {
					fmt.Println(strconv.Itoa(year) + " is leap year.")
				} else {
					fmt.Println(strconv.Itoa(year) + " is not leap year.")
				}
				var horo, zodiac string
				horo = horoscope.GetHoroscope(month, day)
				zodiac = horoscope.GetZodiac(month, day, year)
				fmt.Println("Your horoscope is " + horo + " and your zodiac is " + zodiac)
			case "6":
				run = false
			default:
				fmt.Println("Please enter index between 1~6")
			}
		}
	}
}
