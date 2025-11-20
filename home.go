package main

import (
	"casabaldini/internal/db"
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
type Slider struct {
	ID      int
	Codice  string
	Codice2 string
	Img     string
	Titolo  string
	Caption string
	Link    string
	Testo   string
}

type Submenus struct {
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
	Submenus []Submenus
}

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title        string
		Items        []string
		Year         int
		PageContent  string
		ContentBlock string
		Sliders      []Slider
		Menus        []Menus
		Submenus     []Submenus
	}
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

	data := PageData{
		Title:        "Home Page",
		Items:        []string{"Uno", "Due", "Tre"},
		Year:         time.Now().Year(),
		PageContent:  "home",
		ContentBlock: "content_home",
	}
	Sliders, _ := GetSliders()
	Menus, _ := Menu()
	Submenus, _ := Menu()
	//Submenus, _ := Menu()
	templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"Sliders": Sliders, "Menus": Menus, "Submenus": Submenus, "data": data,
	})
	//if err := templates.ExecuteTemplate(w, "home", data); err != nil {
	//
	//	log.Println("Errore template:", err)
	//	http.Error(w, "Errore interno", 500)

	//}

}

func GetSliders() ([]Slider, error) {

	rows, err := db.DB.Query("SELECT id, codice, codice2, img, titolo, caption, link, testo FROM beb_slider order by codice2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var sliders []Slider

	for rows.Next() {
		var m Slider
		rows.Scan(&m.ID, &m.Codice, &m.Codice2, &m.Img, &m.Titolo, &m.Caption, &m.Link, &m.Testo)

		sliders = append(sliders, m)
	}

	return sliders, nil
}
func Menu() ([]Menus, error) {

	rows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo,link FROM menu WHERE livello=? AND attivo=?", 2, 1)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var menus []Menus
	var submenus []Submenus
	for rows.Next() {
		var m Menus
		rows.Scan(&m.ID, &m.Codice, &m.Radice, &m.Livello, &m.Titolo, &m.Link)

		subRows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo, link FROM submenu WHERE radice = ?", m.Codice)
		if err != nil {

			log.Println(err)
			return nil, err
		}
		defer rows.Close()

		for subRows.Next() {
			var s Submenus
			subRows.Scan(&s.ID, &s.Codice, &s.Radice, &s.Livello, &s.Titolo, &s.Link)
			m.Submenus = append(m.Submenus, s)
			submenus = append(submenus, s)
		}
		defer subRows.Close()

		menus = append(menus, m)

	}
	return menus, nil
}

func Submenu() ([]Submenus, error) {

	rows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo,link FROM menu WHERE livello=? AND attivo=?", 2, 1)

	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var menus []Menus
	var submenus []Submenus
	for rows.Next() {
		var m Menus
		rows.Scan(&m.ID, &m.Codice, &m.Radice, &m.Livello, &m.Titolo, &m.Link)

		subRows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo, link FROM submenu WHERE radice = ?", m.Codice)
		if err != nil {

			log.Println(err)
			return nil, err
		}
		defer rows.Close()

		for subRows.Next() {
			var s Submenus
			subRows.Scan(&s.ID, &s.Codice, &s.Radice, &s.Livello, &s.Titolo, &s.Link)
			m.Submenus = append(m.Submenus, s)
			submenus = append(submenus, s)
		}
		defer subRows.Close()

		menus = append(menus, m)

	}
	return submenus, nil
}
