package DaysToBday

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func CheckError(e error) {
	if e != nil {
		fmt.Println(e)
	}
}

func formatDate(year, month, date int) time.Time {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	return time.Date(year, time.Month(month), date, 0, 0, 0, 0, loc)
}

func getLocalTime() time.Time {
	loc, error := time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc)
	return now
}

func splitBday(bDay string) time.Time {
	splitBirthDay := strings.Split(bDay, "/")
	month, error := strconv.Atoi(splitBirthDay[0])
	CheckError(error)
	day, error := strconv.Atoi(splitBirthDay[1])
	CheckError(error)
	year, error := strconv.Atoi(splitBirthDay[2])
	CheckError(error)

	birthDate := formatDate(year, month, day)
	return birthDate
}
