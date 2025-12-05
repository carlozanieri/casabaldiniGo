package main

import (
	"casabaldiniGo/internal/db"
	"log"
	"net/http"
)

func main() {
	db.Init()
	// Routing semplice
	http.HandleFunc("/", HomeHandler)
	http.HandleFunc("/lecamere", LecamereHandler)
	http.HandleFunc("/lasala", LasalaHandler)
	http.HandleFunc("/ilpaese", IlpaeseHandler)

	// File statici
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	log.Println("Server avviato su http://0.0.0.0:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))
}
