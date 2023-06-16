package controller

import (
	"bytes"
	"html/template"
	"net/http"
	"path/filepath"
	"portfolio/model"
)

/*
Charge le template de l'accueil
*/
func Home(w http.ResponseWriter, r *http.Request) {
	names := make(map[string]string)
	names["owner"] = "damien"

	var listProject = make(map[string]model.Project)
	var project model.Project
	project.Name = "projet2"
	project.Description = "Description du projet"
	project.Technologies = []string{"C#", "Go"}
	listProject["0"] = project

	renderTemplates(w, "home", &model.TemplateData{
		StringData:  names,
		ListProject: listProject,
	})
}

/*
Charge le template "Contact"
*/
func Contact(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "contact", &model.TemplateData{})
}

/*
Charge le template : "à propros"
*/
func About(w http.ResponseWriter, r *http.Request) {
	foo := make(map[string]int)
	foo["owner"] = 30
	renderTemplates(w, "about", &model.TemplateData{
		IntData: foo,
	})
}

/*
Charge un template html
@param http.ResponseWriter ....
@param string template à prendre en compte
*/
func renderTemplates(w http.ResponseWriter, tmplName string, templateData *model.TemplateData) {
	templateCache, err := createTemplateCache() // Charge le cache
	if err != nil {
		panic(err)
	}

	tmpl, ok := templateCache[tmplName+".page.tmpl"]
	if !ok {
		http.Error(w, "le template n'existe pas", http.StatusInternalServerError)
		return
	}

	// Stock tout dans un buffer
	buffer := new(bytes.Buffer)
	tmpl.Execute(buffer, templateData)
	buffer.WriteTo(w) // permet d'écrire les différents données dans le buffer

}

/* Créer un cache permettant de créer un cache avec les templates */
func createTemplateCache() (map[string]*template.Template, error) {
	cache := map[string]*template.Template{}
	var err error

	pages, erreur := filepath.Glob("./templates/*.page.tmpl") // récupère tout selon un paternes dans un tableau
	if erreur != nil {
		err = erreur
	}

	// Parcourt chaque élément du tableau
	for _, page := range pages {
		name := filepath.Base(page)
		tmpl := template.Must(template.ParseFiles(page)) // récupère le path de la page

		// Pareil mais pour chaque layout
		_, erreur := filepath.Glob("./templates/layouts/*.layout.tmpl")
		if erreur != nil {
			err = erreur
		}

		// Récupère le layout et le met dans chacunes des pages
		tmpl.ParseGlob("./templates/layouts/*.layout.tmpl")

		cache[name] = tmpl
	}

	return cache, err
}
