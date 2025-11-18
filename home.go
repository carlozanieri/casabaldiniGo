package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type PageData struct {
	Title        string
	Items        []string
	Year         int
	PageContent  string
	ContentBlock string
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title        string
		Items        []string
		Year         int
		PageContent  string
		ContentBlock string
	}
	data := PageData{
		Title:        "Home Page",
		Items:        []string{"Uno", "Due", "Tre"},
		Year:         time.Now().Year(),
		PageContent:  "home",
		ContentBlock: "content_home",
	}
	//var u PageData
	//var pagina = &u.PageContent
	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
		log.Println("Errore template:", err)
		http.Error(w, "Errore interno", 500)
	}
}
