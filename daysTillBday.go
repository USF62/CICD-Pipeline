package main

import (
	"fmt"
	"time"
  "strings"
	"strconv"
)


func createDate(year, month, date int) time.Time {
	return time.Date(year,time.Month(month), date, 0,0,0,0, time.UTC)
}


func calculateDaysToBday(bDay string){
  splitBirthDay := strings.Split(bDay, "/")
  year, error := strconv.Atoi(splitBirthDay[0])
  month, error := strconv.Atoi(splitBirthDay[1])
  day, error := strconv.Atoi(splitBirthDay[2])

  if(error != nil){
    return
  }
	
	birthDate := createDate(year, month, day) //This creates a Date with user inputted BirthDay
	now := time.Now()  //Current Time

	if birthDate.Month() == now.Month() {

		if birthDate.Day() < now.Day() { //Upcoming birthday next year

			fmt.Println("Your next birthday is next year!")
                	bDayYear := int(now.Year()) +1
                	nextBirthDay := createDate(bDayYear, int(birthDate.Month()), int(birthDate.Day()))
              	 	daysToBday := nextBirthDay.Sub(now).Hours()/24
                	fmt.Printf("You have around %.2f days until your birthday.\n", daysToBday)

		} else { //Upcoming birthday in the same month, same year

			fmt.Println("Your next birthday is this year!")
        		nextBirthDay := createDate(int(now.Year()), int(birthDate.Month()), int(birthDate.Day()))
        		daysToBday := nextBirthDay.Sub(now).Hours()/24
        		fmt.Printf("You have around %.2f days until your birthday.\n", daysToBday)
		}

	} else if birthDate.Month() < now.Month() { //Upcoming birthday next year

		fmt.Println("Your next birthday is next year!")
		dyear := int(now.Year()) +1
		nextBirthDay := createDate(dyear, int(birthDate.Month()), int(birthDate.Day()))
	  daysToBday := nextBirthDay.Sub(now).Hours()/24
		fmt.Printf("You have around %.2f days until your birthday.\n", daysToBday)

	} else { //Upcoming birthday in the same year

		fmt.Println("Your next birthday is this year!")
		nextBirthDay := createDate(int(now.Year()), int(birthDate.Month()), int(birthDate.Day()))
		daysToBday := nextBirthDay.Sub(now).Hours()/24
		fmt.Printf("You have around %.2f days until your birthday.\n", daysToBday)

	}
}

func main() {

	fmt.Println("Hi, Please input your birthday in the following format in numbers:\nyear/month/date  Ex: 2000/03/24")
	var bDay string
	fmt.Scanln(&bDay)

  calculateDaysToBday(bDay)	
}