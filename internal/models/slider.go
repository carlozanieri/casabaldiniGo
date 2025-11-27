package models

import (
	"casabaldiniGo/internal/db"
)

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
