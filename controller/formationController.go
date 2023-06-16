/*
Gère la page sur les formations
*/

package controller

import (
	"net/http"
	"portfolio/model"
)

func Formation(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "formation", &model.TemplateData{})
}
