/*
Gère la page à propos de moi
*/

package controller

import (
	"net/http"
	"portfolio/model"
)

func Me(w http.ResponseWriter, r *http.Request) {
	renderTemplates(w, "me", &model.TemplateData{})
}
