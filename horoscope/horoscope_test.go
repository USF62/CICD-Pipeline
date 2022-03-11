package horoscope

import "testing"

func TestHoroscope(t *testing.T) {
	if getHoroscope(1, 19) != "Capricorn" {
		t.Errorf("Expected Capricorn")
	}
}

func TestTableCalculate(t *testing.T) {
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
			t.Errorf("Test Failed: %d, expected: %q, received: %q", test.inputID, test.expected, output)
		}
	}
}
