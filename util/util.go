package util

import (
	"fmt"
	"time"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func FormatDate(year, month, date int) time.Time {
	loc, err := time.LoadLocation("America/Los_Angeles")
	CheckError(err)
	return time.Date(year, time.Month(month), date, 0, 0, 0, 0, loc)
}

func GetLocalTime() time.Time {
	loc, err := time.LoadLocation("America/Los_Angeles")
	CheckError(err)
	now := time.Now().UTC().In(loc)
	return now
}

//Not using function
//func SplitBday(bDay string) time.Time {
//	splitBirthDay := strings.Split(bDay, "/")
//	month, err := strconv.Atoi(splitBirthDay[0])
//	CheckError(err)
//	day, err := strconv.Atoi(splitBirthDay[1])
//	CheckError(err)
//	year, err := strconv.Atoi(splitBirthDay[2])
//	CheckError(err)
//
//	birthDate := FormatDate(year, month, day)
//	return birthDate
//}

//Not using function
//func HandleInput(s, sep string) []string {
//	repeat := true
//	var final []string
//
//	for repeat == true {
//		final = strings.Split(s, sep)
//		if len(final) == 1 {
//			fmt.Printf("Input incorrect! Please use %s as the seperators. \n", sep)
//			_, err := fmt.Scanln(&s)
//			if err != nil {
//				return nil
//			}
//		} else {
//			repeat = false
//		}
//	}
//	return final
//}

func Is30DayMonth(month int) bool {
	return month == 4 || month == 6 || month == 9 || month == 11
}
