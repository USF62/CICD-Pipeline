package main

import (
	"fmt"
	"github.com/venkatrajm815/USF62-Pipeline/calculate"
	"strconv"
	"time"
)

func main() {
	fmt.Println("enter date of birth (MM/DD/YYYY): ")
	todayDate := time.Now().Format("01/02/2006")
	var birthDate string
	fmt.Scanln(&birthDate)

	test := timecalculator.HandleInput(todayDate, "/")
	month, _ := strconv.Atoi(test[0])
	day, _ := strconv.Atoi(test[1])
	year, _ := strconv.Atoi(test[2])

	test = calculate.HandleInput(birthDate, "/")
	monthBirth, _ := strconv.Atoi(test[0])
	dayBirth, _ := strconv.Atoi(test[1])
	yearBirth, _ := strconv.Atoi(test[2])

	monthDiff, dayDiff, yearDiff := timecalculator.Calculation(month, day, year, monthBirth, dayBirth, yearBirth)

	fmt.Printf("You are %d years, %d months, %d days old.", yearDiff, monthDiff, dayDiff)

}
