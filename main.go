package main

import (
	"log"
	"net/http"
)

func main() {

	// Routing semplice
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/about", AboutHandler)
	http.HandleFunc("/contact", ContactHandler)

	// File statici
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server avviato su http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
