package handlers

import (
	"casabaldini/internal/models"
	"html/template"
	"log"
	"net/http"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func Home(w http.ResponseWriter, r *http.Request) {

	sliders, err := models.GetSliders()
	if err != nil {
		log.Println(err)
		http.Error(w, "Errore interno", 500)
		return
	}

	err = templates.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Sliders": sliders,
	})
	if err != nil {
		log.Println(err)
	}
}
