package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "home")
}

func about(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "about")
}

func project(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "project")
}

func contact(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "contact")
}

// w =
// tmpl = template de base
func renderTemplates(w http.ResponseWriter, tmpl string) {
	t, err := template.ParseFiles("./templates/" + tmpl + ".page.tmpl")

	if err != nil {
		fmt.Println("error")
	}
	t.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/project", project)
	http.HandleFunc("/contact", contact)
	http.ListenAndServe(":5500", nil)
}
