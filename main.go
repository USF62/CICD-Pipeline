package main

import (
	"fmt"
	"USF62-Pipeline/calculate"
	"USF62-Pipeline/util"
)

func main() {
	fmt.Println("enter date of birth (MM/DD/YYYY): ")
	var birthDate string
	for true {
		fmt.Scanln(&birthDate)
		if err := util.VerifyInput(birthDate); err == nil {
			break
		}
	}

	date, err := util.GetDate(birthDate)
	if err != nil {
		fmt.Println("Error getting date:", err)
	}

	yearDiff, monthDiff, dayDiff := calculate.CalculationSince(date)

	fmt.Printf("You are %d years, %d months, %d days old.", yearDiff, monthDiff, dayDiff)
}
