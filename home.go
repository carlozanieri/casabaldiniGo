package main

import (
	"casabaldini/internal/db"
	"html/template"
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

var templates = template.Must(template.ParseGlob("templates/*.html"))

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	type PageData struct {
		Title        string
		Items        []string
		Year         int
		PageContent  string
		ContentBlock string
		Sliders      []Slider
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
	templates.ExecuteTemplate(w, "home", map[string]interface{}{
		"Sliders": Sliders, "data": data,
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
