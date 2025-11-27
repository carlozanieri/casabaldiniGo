package models

import (
	"casabaldiniGo/internal/db"
)

type Submenu struct {
	ID      int
	Codice  string
	Radice  string
	Livello int
	Titolo  string
	Link    string
}

type Menu struct {
	ID       int
	Codice   string
	Radice   string
	Livello  int
	Titolo   string
	Link     string
	Submenus []Submenu
}

type Link struct {
	ID     int
	Codice string
	Titolo string
	Link   string
}

func GetMenus() ([]Menu, error) {

	rows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo,link FROM menu WHERE livello=? AND attivo=?", 2, 1)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var menus []Menu

	for rows.Next() {
		var m Menu
		rows.Scan(&m.ID, &m.Codice, &m.Radice, &m.Livello, &m.Titolo, &m.Link)

		subRows, err := db.DB.Query("SELECT id, codice,  radice, livello, titolo, link FROM submenu WHERE radice = ?", m.Codice)
		if err != nil {
			return nil, err
		}

		for subRows.Next() {
			var s Submenu
			subRows.Scan(&s.ID, &s.Codice, &s.Radice, &s.Livello, &s.Titolo, &s.Link)
			m.Submenus = append(m.Submenus, s)
		}
		subRows.Close()

		menus = append(menus, m)
	}

	return menus, nil
}
