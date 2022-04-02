package webapp

import (
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

var templates *template.Template

func Start() {
	templates = template.Must(template.ParseGlob("webapp/templates/*.html"))
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler).Methods("GET")
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
	//r.HandleFunc("/results", ResultHandler).Methods("POST")

	log.Fatal(http.ListenAndServe("localhost:8080", r))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "home.html", nil)
	if err != nil {
		return
	}
}

//// Save given page's body to a new text file.
//func (p *Page) Save() error {
//	filename := p.Title + ".txt"
//	return os.WriteFile(filename, p.Body, 0600)
//}
//
//// LoadPage Reads file contents into a new body, returns nil if page is successfully loaded
//func LoadPage(title string) (*Page, error) {
//	filename := title + ".html"
//	body, err := os.ReadFile(filename)
//	if err != nil {
//		return nil, err
//	}
//	return &Page{Title: title, Body: body}, nil
//}
//
//// ViewHandler allows users to view a page and handles URLs prefixed with "/view/"
//// also handles non-existing pages by redirecting client back to the edit page
//func ViewHandler(w http.ResponseWriter, r *http.Request) {
//	p, err := LoadPage(title)
//	if err != nil {
//		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
//		return
//	}
//	renderTemplate(w, "view", p)
//}
//
//// EditHandler Executes template and write generated HTML
//func EditHandler(w http.ResponseWriter, r *http.Request) {
//	p, err := LoadPage(title)
//	if err != nil {
//		p = &Page{Title: title}
//	}
//	renderTemplate(w, "edit", p)
//}
//
//// SaveHandler handles submission of forms
//func SaveHandler(w http.ResponseWriter, r *http.Request) {
//	body := r.FormValue("body")
//	p := &Page{Title: title, Body: []byte(body)}
//	err := p.Save()
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//		return
//	}
//	http.Redirect(w, r, "/view/"+title, http.StatusFound)
//}
//
//// Generates HTML template
//func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
//	err := templates.ExecuteTemplate(w, tmpl+".html", p)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusInternalServerError)
//	}
//}
//
//var validPath = regexp.MustCompile("^/(edit|save|view)/([a-zA-Z0-9]+)$")
//
//// MakeHandler Checks if the request is valid and executes request
//func MakeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		m := validPath.FindStringSubmatch(r.URL.Path)
//		if m == nil {
//			http.NotFound(w, r)
//			return
//		}
//		fn(w, r, m[2])
//	}
//}
