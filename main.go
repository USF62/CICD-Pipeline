package main

import (
	"fmt"
	"github.com/venkatrajm815/USF62-Pipeline/DaysToBday"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi, Please input your birthday in the following format in numbers:\nmonth/date/year  Ex: 03/24/2000")
	var bDay string
	fmt.Scanln(&bDay)
	daysToBirthday := DaysToBday.CalculateDaysToBday(DaysToBday.SplitBday(bDay))
	splitBirthDay := strings.Split(bDay, "/")
	month, error := strconv.Atoi(splitBirthDay[0])
	DaysToBday.CheckError(error)
	day, error := strconv.Atoi(splitBirthDay[1])
	DaysToBday.CheckError(error)
	year, error := strconv.Atoi(splitBirthDay[2])
	DaysToBday.CheckError(error)

	horoscope := DaysToBday.GetHoroscope(month, day)
	zodiac := DaysToBday.GetZodiac(month, day, year)

	fmt.Println(daysToBirthday)
	fmt.Println(horoscope)
	fmt.Println(zodiac)
}
