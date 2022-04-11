package webapp

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

var client *redis.Client
var templates *template.Template

type Birthdate struct {
	Month int
	Day   int
	Year  int
}

func Start() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:8080",
	})
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

	userBD := strings.Split(r.FormValue("birthdate"), "-")
	userYear, _ := strconv.Atoi(userBD[0])
	userMonth, _ := strconv.Atoi(userBD[1])
	userDay, _ := strconv.Atoi(userBD[2])

	birthdate := Birthdate{
		Month: userMonth,
		Day:   userDay,
		Year:  userYear,
	}
	w.Write([]byte(userBD[0]))
	fmt.Print(birthdate)
}
