package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

/*
Charge un template html
@param http.ResponseWriter ....
@param string template à prendre en compte
*/
func renderTemplates(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")

	if err != nil {
		fmt.Println("error")
	}
	t.Execute(w, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about")
}

func Project(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "project")
}

func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "contact")
}
