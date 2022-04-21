package webapp

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/venkatrajm815/USF62-Pipeline/calculate"
	"github.com/venkatrajm815/USF62-Pipeline/daysToBday"
	"github.com/venkatrajm815/USF62-Pipeline/horoscope"
	"github.com/venkatrajm815/USF62-Pipeline/leapYear"
	"github.com/venkatrajm815/USF62-Pipeline/util"
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

func Start() {
	templates = template.Must(template.ParseFiles("webapp/templates/home.html", "webapp/templates/results.html"))
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

// GET handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		fmt.Print(err)
	}
}

// POST handler
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

func initBirthDate(birthdate []string) {
	userBD.Year, _ = strconv.Atoi(birthdate[0])
	userBD.Month, _ = strconv.Atoi(birthdate[1])
	userBD.Day, _ = strconv.Atoi(birthdate[2])
	userBD.Birthdate = util.FormatDate(userBD.Year, userBD.Month, userBD.Day)
}

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
