package DaysToBday

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hi, Please input your birthday in the following format in numbers:\nmonth/date/year  Ex: 03/24/2000")
	var bDay string
	fmt.Scanln(&bDay)
	daysToBirthday := CalculateDaysToBday(splitBday(bDay))
	splitBirthDay := strings.Split(bDay, "/")
	month, error := strconv.Atoi(splitBirthDay[0])
	CheckError(error)
	day, error := strconv.Atoi(splitBirthDay[1])
	CheckError(error)
	year, error := strconv.Atoi(splitBirthDay[2])
	CheckError(error)

	horoscope := getHoroscope(month, day)
	zodiac := getZodiac(month, day, year)

	fmt.Println(daysToBirthday)
	fmt.Println(horoscope)
	fmt.Println(zodiac)
}
