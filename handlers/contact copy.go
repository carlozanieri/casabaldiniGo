package main

import (
	"log"
	"net/http"
	"time"
)

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title       string
		Items       []string
		Year        int
		PageContent string // <- aggiungi questo
	}

	data := PageData{
		Title:       "Contact",
		Year:        time.Now().Year(),
		PageContent: "contact",
	}

	if err := templates.ExecuteTemplate(w, "contact", data); err != nil {
		log.Println("Errore template Contact:", err)
		http.Error(w, "Errore interno", 500)
	}
}
