package calculate

import (
	"fmt"
	"testing"
	"time"
)

func TestCalc(t *testing.T) {
	monthTest, dayTest, yearTest := Calculation(
		time.Date(2000, 9, 28, 0, 0, 0, 0, time.UTC),
		time.Date(2022, 3, 19, 0, 0, 0, 0, time.UTC))

	if monthTest == 5 && dayTest == 32 && yearTest == 21 {
		return
	} else {
		fmt.Println(monthTest, dayTest, yearTest)
		t.Fatal("Error! Calculated return and actual does not match up!")
	}
}
