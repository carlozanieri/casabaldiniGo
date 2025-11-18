package handlers

import (
	"casabaldini/internal/db"
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

type Slider struct {
	ID         int
	Codice     string
	Codice2    string
	Img        string
	Titolo     string
	Caption    string
	Link       string
	Testo      string
	StaticPath string
}

type Link struct {
	ID     int
	Codice string
	Titolo string
	Link   string
}

func Home(w http.ResponseWriter, r *http.Request) {
	templates := template.Must(template.ParseGlob(filepath.Join("templates", "*.html")))
	rows, err := db.DB.Query("SELECT id, codice, codice2, img, titolo, caption, link, testo FROM beb_slider")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var sliders []Slider

	for rows.Next() {
		var u Slider
		//u.StaticPath = "/static"
		if err := rows.Scan(&u.ID, &u.Codice, &u.Codice2, &u.Img, &u.Titolo, &u.Caption, &u.Link, &u.Testo); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			u.StaticPath = "/static"
			return
		}
		sliders = append(sliders, u)
	}

	err = templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"Sliders": sliders,
	})

	if err != nil {
		log.Println(err)
	}
}
