package calculate

import (
	"fmt"
	"testing"
)

func testCalc(t *testing.T){
	monthBirth := 9
	dayBirth := 28
	yearBirth := 2000

	month := 3
	day := 19
	year := 2022

	monthTest, dayTest, yearTest := Calculation(month, day, year, monthBirth, dayBirth, yearBirth)

	if monthTest == 6 && dayTest == 9 && yearTest == 22 {
		return
	} else {
		t.ErrorF("Error!")
	}
}
