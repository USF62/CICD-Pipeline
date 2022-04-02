package horoscope

import "testing"

func TestHoroscope(t *testing.T) {
	if getHoroscope(1, 19) != "Capricorn" {
		t.Errorf("Expected Capricorn")
	}
}

func TestTableHoroscope(t *testing.T) {
	var tests = []struct {
		inputID    int
		inputMonth int
		inputDay   int
		expected   string
	}{
		{1, 12, 22, "Capricorn"},
		{2, 1, 19, "Capricorn"},
		{3, 1, 20, "Aquarius"},
		{4, 2, 18, "Aquarius"},
		{5, 2, 19, "Pisces"},
		{6, 3, 20, "Pisces"},
		{7, 3, 21, "Aries"},
		{8, 4, 19, "Aries"},
		{9, 4, 20, "Taurus"},
		{10, 5, 20, "Taurus"},
		{11, 5, 21, "Gemini"},
		{12, 6, 20, "Gemini"},
		{13, 6, 21, "Cancer"},
		{14, 7, 22, "Cancer"},
		{15, 7, 23, "Leo"},
		{16, 8, 22, "Leo"},
		{17, 8, 23, "Virgo"},
		{18, 9, 22, "Virgo"},
		{19, 9, 23, "Libra"},
		{20, 10, 22, "Libra"},
		{21, 10, 23, "Scorpio"},
		{22, 11, 21, "Scorpio"},
		{23, 11, 22, "Sagittarius"},
		{24, 12, 21, "Sagittarius"},
		{25, 2, 29, "Pisces"},
	}
	for _, test := range tests {
		if output := getHoroscope(test.inputMonth, test.inputDay); output != test.expected {
			t.Errorf("Test Failed: %d, Expected: %q, Received: %q", test.inputID, test.expected, output)
		}
	}
}

func TestZodiac(t *testing.T) {
	if getZodiac(2, 3, 2000) != "Rabbit" {
		t.Errorf("Expected Rabbit")
	}
}

func TestTableZodiac(t *testing.T) {
	var tests = []struct {
		inputID    int
		inputMonth int
		inputDay   int
		inputYear  int
		expected   string
	}{
		{1, 2, 3, 2000, "Rabbit"},
		{2, 2, 4, 2000, "Dragon"},
		{3, 2, 3, 2001, "Dragon"},
		{4, 2, 4, 2001, "Snake"},
		{5, 2, 3, 2002, "Snake"},
		{6, 2, 4, 2002, "Horse"},
		{7, 2, 3, 2003, "Horse"},
		{8, 2, 4, 2003, "Sheep"},
		{9, 2, 3, 2004, "Sheep"},
		{10, 2, 4, 2004, "Monkey"},
		{11, 2, 3, 2005, "Monkey"},
		{12, 2, 4, 2005, "Rooster"},
		{13, 2, 3, 2006, "Rooster"},
		{14, 2, 4, 2006, "Dog"},
		{15, 2, 3, 2007, "Dog"},
		{16, 2, 4, 2007, "Pig"},
		{17, 2, 3, 2008, "Pig"},
		{18, 2, 4, 2008, "Rat"},
		{19, 2, 3, 2009, "Rat"},
		{20, 2, 4, 2009, "Ox"},
		{21, 2, 3, 2010, "Ox"},
		{22, 2, 4, 2010, "Tiger"},
		{23, 2, 3, 2011, "Tiger"},
		{24, 2, 4, 2011, "Rabbit"},
		{25, 2, 3, 2012, "Rabbit"},
		{26, 2, 4, 2012, "Dragon"},
	}
	for _, test := range tests {
		if output := getZodiac(test.inputMonth, test.inputDay, test.inputYear); output != test.expected {
			t.Errorf("Test Failed: %d, Expected: %q, Received: %q", test.inputID, test.expected, output)
		}
	}
}
