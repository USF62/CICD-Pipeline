package main

import (
	"fmt"
	"time"
  "strings"
	"strconv"
)


func formatDate(year, month, date int) time.Time {
	return time.Date(year,time.Month(month), date, 0,0,0,0, time.UTC)
}

func CheckError( e error){
	if e != nil{
		fmt.Println(e)
	}
}


func calculateDaysToBday(bDay string){
  splitBirthDay := strings.Split(bDay, "/")
  year, error := strconv.Atoi(splitBirthDay[0])
  month, error := strconv.Atoi(splitBirthDay[1])
  day, error := strconv.Atoi(splitBirthDay[2])
  CheckError(error)
	
	birthDate := formatDate(year, month, day) //This creates a Date with user inputted BirthDay
	loc, error :=  time.LoadLocation("America/Los_Angeles")
	CheckError(error)
	now := time.Now().UTC().In(loc) 
	
	if ((birthDate.Month() == now.Month() && birthDate.Day() >= now.Day() ) || (birthDate.Month() > now.Month())){
		if(birthDate.Day() == now.Day()){
			fmt.Println("Today is your birthday. Happy Birthday!")
			return
		}
		fmt.Println("Your next birthday is this year!")
	  nextBirthDay := formatDate(int(now.Year()), int(birthDate.Month()), int(birthDate.Day()))
	  daysToBday := nextBirthDay.Sub(now).Hours()/24
	  fmt.Printf("You have around %.2f days until your birthday.\n", daysToBday)
	} else if ((birthDate.Month() == now.Month() && birthDate.Day() < now.Day() ) || (birthDate.Month() < now.Month())){
		fmt.Println("Your next birthday is next year!")
		bDayYear := int(now.Year()) +1
		nextBirthDay := formatDate(bDayYear, int(birthDate.Month()), int(birthDate.Day()))
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