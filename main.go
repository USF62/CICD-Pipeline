package main

import (
	"github.com/venkatrajm815/USF62-Pipeline/webapp"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("webapp/view/", webapp.MakeHandler(webapp.ViewHandler))
	http.HandleFunc("webapp/edit/", webapp.MakeHandler(webapp.EditHandler))
	http.HandleFunc("webapp/save/", webapp.MakeHandler(webapp.SaveHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))
}
