package calculate

import (
	"time"
	"fmt"
)

func CalculationSince(date time.Time) (int, int, int) {
	return Calculation(time.Now().UTC(), date)
}

func Calculation(today, birthday time.Time) (int, int, int) {
	fmt.Println("Todays date:", today)
	fmt.Println("Parsed date:", birthday)

	if today.Location() != birthday.Location() {
		birthday = birthday.In(today.Location())
	}

	y1, M1, d1 := birthday.Date()
	y2, M2, d2 := today.Date()

	year := int(y2 - y1)
	month := int(M2 - M1)
	day := int(d2 - d1)

	// Normalize negative values
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}

	if month < 0 {
 		month += 12
		year--
	}

	return year, month, day
}
