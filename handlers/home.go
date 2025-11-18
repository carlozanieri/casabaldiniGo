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
	ContentBlock string // not mandatory with base_top/base_bottom approach
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{
		Title: "Home",
		Items: []string{"Uno", "Due", "Tre"},
		Year:  time.Now().Year(),
	}
	if err := templates.ExecuteTemplate(w, "home", data); err != nil {
		log.Println("template error home:", err)
		http.Error(w, "internal error", 500)
	}
}

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "About", Year: time.Now().Year()}
	if err := templates.ExecuteTemplate(w, "about", data); err != nil {
		log.Println("template error about:", err)
		http.Error(w, "internal error", 500)
	}
}

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	data := PageData{Title: "Contact", Year: time.Now().Year()}
	if err := templates.ExecuteTemplate(w, "contact", data); err != nil {
		log.Println("template error contact:", err)
		http.Error(w, "internal error", 500)
	}
}
