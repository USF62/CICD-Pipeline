package DaysToBday

var horoscope = []string{"Capricorn", "Aquarius", "Pisces", "Aries", "Taurus", "Gemini",
	"Cancer", "Leo", "Virgo", "Libra", "Scorpio", "Sagittarius"}

var zodiac = []string{"Monkey", "Rooster", "Dog", "Pig", "Rat", "Ox",
	"Tiger", "Rabbit", "Dragon", "Snake", "Horse", "Sheep"}

/* Returns horoscope sign based on the given birthday.
 * Assumes given inputs are positive. */
func getHoroscope(month int, day int) string {
	if (month == 1 || month == 4) && day <= 19 ||
		(month == 2 && day <= 18) ||
		(month == 3 || month == 5 || month == 6) && day <= 20 ||
		((month == 11 || month == 12) && day <= 21) ||
		(month >= 7 && month <= 10) && day <= 22 {
		return horoscope[month-1]
	}
	if month == 12 {
		return horoscope[0]
	}
	return horoscope[month]
}

/* Checks if the birthday will be the previous year's zodiac sign
 * according to the Chinese calendar. (Rough estimate) */
func isPrevZodiac(month int, day int) bool {
	return month == 2 && day < 4
}

/* Returns zodiac sign based on the given birthdate.
 * Assumes given inputs are positive. */
func getZodiac(month int, day int, year int) string {
	year = year % 12
	if isPrevZodiac(month, day) {
		if year == 0 {
			year = 11
		} else {
			year--
		}
	}
	return zodiac[year]
}
