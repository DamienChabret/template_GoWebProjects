/*
Gère la page sur les expériences
*/

package controller

import (
	"net/http"
	"portfolio/model"
)

func Experience(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "experience", &model.TemplateData{})
}
