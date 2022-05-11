package webapp

import (
	"fmt"
	"github.com/USF62/CICD-Pipeline/calculate"
	"github.com/USF62/CICD-Pipeline/daysToBday"
	"github.com/USF62/CICD-Pipeline/horoscope"
	"github.com/USF62/CICD-Pipeline/leapYear"
	"github.com/USF62/CICD-Pipeline/util"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

type Birthdate struct {
	Month, Day, Year int
	Birthdate        time.Time
}

var templates *template.Template
var userBD = Birthdate{}

// Start reads necessary HTML files and sets up router and port
func Start() {
	templates = template.Must(template.ParseFiles("webapp/templates/home.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/", resultsHandler).Methods("POST")
	http.Handle("/", r)

	var port = os.Getenv("PORT")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		fmt.Print(err)
	}
	log.Fatal(http.ListenAndServe(":"+port, r))
}

// homeHandler loads home HTML file
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		fmt.Print(err)
	}
}

// resultsHandler parses user input and rewrites web page with results
func resultsHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		return
	}
	if err := r.ParseForm(); err != nil {
		fmt.Print(err)
		return
	}
	birthdate := strings.Split(r.FormValue("birthdate"), "-")
	initBirthDate(birthdate)
	results := writeResults()
	fmt.Print(results)
	w.Write([]byte(results))
}

// initBirthdate parses birthdate into separate variables
func initBirthDate(birthdate []string) {
	userBD.Year, _ = strconv.Atoi(birthdate[0])
	userBD.Month, _ = strconv.Atoi(birthdate[1])
	userBD.Day, _ = strconv.Atoi(birthdate[2])
	userBD.Birthdate = util.FormatDate(userBD.Year, userBD.Month, userBD.Day)
}

// writeResults writes and returns a results string of
// total age, days until birthday,horoscope/zodiac sign, and leap year validity
func writeResults() string {
	builder := strings.Builder{}
	month, day, year := calculate.TotalAge(userBD.Birthdate, time.Now())
	totalAge := fmt.Sprintf("Your total age is: %d year(s) %d month(s) %d day(s)\n", year, month, day)
	builder.Write([]byte(totalAge))
	builder.WriteString(daysToBday.CalculateDaysToBdayStr(userBD.Birthdate) + "\n")
	builder.WriteString("Horoscope sign: " + horoscope.GetHoroscope(userBD.Month, userBD.Day) + "\tZodiac Sign: " + horoscope.GetZodiac(userBD.Month, userBD.Day, userBD.Year) + "\n")
	builder.WriteString("Is your birthday on a leap year?: " + strconv.FormatBool(leapYear.IsLeapYear(userBD.Year)))
	return builder.String()
}
