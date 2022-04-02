package webapp

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var client *redis.Client
var templates *template.Template

func Start() {
	client = redis.NewClient(&redis.Options{
		Addr: "localhost:8080",
	})
	templates = template.Must(template.ParseFiles("webapp/templates/home.html"))

	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/", resultsHandler).Methods("POST")
	http.Handle("/", r)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Print(err)
	}
	log.Fatal(http.ListenAndServe("localhost:8080", r))
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
	err := r.ParseForm()
	if err != nil {
		fmt.Print(err)
	}
	birthdate := r.PostFormValue("birthdate")
	client.LPush("birthdate", birthdate)
	http.Redirect(w, r, "/", 302)
}
