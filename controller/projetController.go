/*
Gère la page sur les projets
*/

package controller

import (
	"net/http"
	"portfolio/model"
)

func Project(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "project", &model.TemplateData{})
}
