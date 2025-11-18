package main

import (
	"log"
	"net/http"
	"time"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title       string
		Items       []string
		Year        int
		PageContent string // <- aggiungi questo
	}
	data := PageData{
		Title:       "About",
		Year:        time.Now().Year(),
		PageContent: "about",
	}

	if err := templates.ExecuteTemplate(w, "about", data); err != nil {
		log.Println("Errore template About:", err)
		http.Error(w, "Errore interno", 500)
	}
}
