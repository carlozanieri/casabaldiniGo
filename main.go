package main

import (
	"casabaldini/internal/db"
	"log"
	"net/http"
)

func main() {
	db.Init()
	// Routing semplice
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)

	// File statici
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server avviato su http://localhost:6060")
	log.Fatal(http.ListenAndServe("0.0.0.0:6060", nil))
}
