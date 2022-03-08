package main

func isBetween1and19(day int) bool {
	if day >= 1 && day <= 19 {
		return true
	}
	return false
}

func isBetween1and20(day int) bool {
	if day >= 1 && day <= 20 {
		return true
	}
	return false
}

func isBetween1and21(day int) bool {
	if day >= 1 && day <= 20 {
		return true
	}
	return false
}

func isBetween1and22(day int) bool {
	if day >= 1 && day <= 20 {
		return true
	}
	return false
}

func isBetween21and30(day int) bool {
	if day >= 21 && day <= 30 {
		return true
	}
	return false
}

func isBetween21and31(day int) bool {
	if day >= 21 && day <= 31 {
		return true
	}
	return false
}

func isBetween22and30(day int) bool {
	if day >= 22 && day <= 30 {
		return true
	}
	return false
}

func isBetween22and31(day int) bool {
	if day >= 22 && day <= 31 {
		return true
	}
	return false
}

func isBetween23and31(day int) bool {
	if day >= 23 && day <= 31 {
		return true
	}
	return false
}

func horoscope(month int, day int) string {
	var horoscope string

	if month == 1 {
		if isBetween1and19(day) {
			horoscope = "Capricorn"
		} else if day >= 20 && day <= 31 {
			horoscope = "Aquarius"
		}
	} else if month == 2 {
		if day >= 1 && day <= 18 {
			horoscope = "Aquarius"
		} else if day >= 19 && day <= 29 {
			horoscope = "Pisces"
		}
	} else if month == 3 {
		if isBetween1and20(day) {
			horoscope = "Pisces"
		} else if isBetween21and31(day) {
			horoscope = "Aries"
		}
	} else if month == 4 {
		if isBetween1and20(day) {
			horoscope = "Aries"
		} else if isBetween21and31(day) {
			horoscope = "Taurus"
		}
	} else if month == 5 {
		if isBetween1and20(day) {
			horoscope = "Taurus"
		} else if isBetween21and31(day) {
			horoscope = "Gemini"
		}
	} else if month == 6 {
		if isBetween1and20(day) {
			horoscope = "Gemini"
		} else if isBetween21and30(day) {
			horoscope = "Cancer"
		}
	} else if month == 7 {
		if isBetween1and22(day) {
			horoscope = "Cancer"
		} else if isBetween23and31(day) {
			horoscope = "Leo"
		}
	} else if month == 8 {
		if isBetween1and22(day) {
			horoscope = "Leo"
		} else if isBetween23and31(day) {
			horoscope = "Virgo"
		}
	} else if month == 9 {
		if isBetween1and22(day) {
			horoscope = "Virgo"
		} else if day >= 23 && day <= 30 {
			horoscope = "Libra"
		}
	} else if month == 10 {
		if isBetween1and22(day) {
			horoscope = "Libra"
		} else if isBetween23and31(day) {
			horoscope = "Scorpio"
		}
	} else if month == 11 {
		if isBetween1and21(day) {
			horoscope = "Scorpio"
		} else if isBetween22and30(day) {
			horoscope = "Sagittarius"
		}
	} else if month == 12 {
		if isBetween1and21(day) {
			horoscope = "Sagittarius"
		} else if isBetween22and31(day) {
			horoscope = "Capricorn"
		}
	}
	return horoscope
}

func isPrevZodiac(month int, day int) bool {
	if month == 2 && day < 4 {
		return true
	}
	return false
}

func zodiac(month int, day int, year int) string {
	var zodiac string
	year = year % 12

	if isPrevZodiac(month, day) {
		if year == 0 {
			year = 11
		} else {
			year--
		}
	}

	switch year {
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
