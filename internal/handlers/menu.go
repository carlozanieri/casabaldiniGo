package handlers

import (
	"casabaldini/internal/db"

	"html/template"
	"log"
	"net/http"
)

type Submenu struct {
	ID      int
	Codice  string
	Radice  string
	Livello int
	Titolo  string
	Link    string
}

type Menus struct {
	ID       int
	Codice   string
	Radice   string
	Livello  int
	Titolo   string
	Link     string
	Submenus []Submenu
}

var tmps = template.Must(template.ParseGlob("templates/*.html"))

func Menu(w http.ResponseWriter, r *http.Request) {

	rows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo,link FROM menu WHERE livello=? AND attivo=?", 2, 1)

	if err != nil {
		return
	}
	defer rows.Close()
	var menus []Menus

	for rows.Next() {
		var m Menus
		rows.Scan(&m.ID, &m.Codice, &m.Radice, &m.Livello, &m.Titolo, &m.Link)

		subRows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo, link FROM submenu WHERE radice = ?", m.Codice)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()
		for subRows.Next() {
			var s Submenu
			subRows.Scan(&s.ID, &s.Codice, &s.Radice, &s.Livello, &s.Titolo, &s.Link)
			m.Submenus = append(m.Submenus, s)
		}
		subRows.Close()

		menus = append(menus, m)
	}

	if err != nil {
		log.Println(err)
		http.Error(w, "Errore interno", 500)
		return
	}

	err = tmps.ExecuteTemplate(w, "layout.html", map[string]interface{}{
		"Menu": menus,
	})
	if err != nil {
		log.Println(err)
	}
}
