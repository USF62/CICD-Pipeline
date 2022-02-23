package main

import (
	"fmt"
	"strings"
)

func horoscope(month string, day int) string {
	month = strings.ToLower(month)
	var horoscope string

	if month == "january" {
		if day >= 1 && day <= 19 {
			horoscope = "Capricorn"
		} else if day >= 20 && day <= 30 {
			horoscope = "Aquarius"
		}
	} else if month == "february" {
		if day >= 1 && day <= 18 {
			horoscope = "Aquarius"
		} else if day >= 19 && day <= 29 {
			horoscope = "Pisces"
		}
	} else if month == "march" {
		if day >= 1 && day <= 20 {
			horoscope = "Pisces"
		} else if day >= 21 && day <= 31 {
			horoscope = "Aries"
		}
	} else if month == "april" {
		if day >= 1 && day <= 20 {
			horoscope = "Aries"
		} else if day >= 21 && day <= 30 {
			horoscope = "Taurus"
		}
	} else if month == "may" {
		if day >= 1 && day <= 20 {
			horoscope = "Taurus"
		} else if day >= 21 && day <= 31 {
			horoscope = "Gemini"
		}
	} else if month == "june" {
		if day >= 1 && day <= 20 {
			horoscope = "Gemini"
		} else if day >= 21 && day <= 30 {
			horoscope = "Cancer"
		}
	} else if month == "july" {
		if day >= 1 && day <= 22 {
			horoscope = "Cancer"
		} else if day >= 23 && day <= 31 {
			horoscope = "Leo"
		}
	} else if month == "august" {
		if day >= 1 && day <= 22 {
			horoscope = "Leo"
		} else if day >= 23 && day <= 31 {
			horoscope = "Virgo"
		}
	} else if month == "september" {
		if day >= 1 && day <= 22 {
			horoscope = "Virgo"
		} else if day >= 23 && day <= 30 {
			horoscope = "Libra"
		}
	} else if month == "october" {
		if day >= 1 && day <= 22 {
			horoscope = "Libra"
		} else if day >= 23 && day <= 31 {
			horoscope = "Scorpio"
		}
	} else if month == "november" {
		if day >= 1 && day <= 21 {
			horoscope = "Scorpio"
		} else if day >= 22 && day <= 30 {
			horoscope = "Sagittarius"
		}
	} else if month == "december" {
		if day >= 1 && day <= 21 {
			horoscope = "Sagittarius"
		} else if day >= 22 && day <= 31 {
			horoscope = "Capricorn"
		}
	}
	return horoscope
}

func zodiac(year int) string {
	var zodiac string
	switch year % 12 {
	case 0:
		zodiac = "Monkey"
	case 1:
		zodiac = "Rooster"
	case 2:
		zodiac = "Dog"
	case 3:
		zodiac = "Pig"
	case 4:
		zodiac = "Rat"
	case 5:
		zodiac = "Ox"
	case 6:
		zodiac = "Tiger"
	case 7:
		zodiac = "Rabbit"
	case 8:
		zodiac = "Dragon"
	case 9:
		zodiac = "Snake"
	case 10:
		zodiac = "Horse"
	case 11:
		zodiac = "Sheep"
	}
	return zodiac
}

func main() {
	fmt.Println(horoscope("dEcember", 24))
	fmt.Println(zodiac(2000))
}
