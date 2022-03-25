package calculate

import (
	"testing"
)

func TestCalc(t *testing.T) {
	monthBirth := 9
	dayBirth := 28
	yearBirth := 2000

	month := 3
	day := 19
	year := 2022

	monthTest, dayTest, yearTest := Calculation(month, day, year, monthBirth, dayBirth, yearBirth)

	if monthTest == 5 && dayTest == 32 && yearTest == 21 {
		return
	} else {
		t.Fatal("Error! Calculated return and actual does not match up!")
	}
}
