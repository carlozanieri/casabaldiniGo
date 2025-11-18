package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"
)

// basic LoginHandler: GET shows form, POST checks credentials from DB
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		templates.ExecuteTemplate(w, "login", nil)
		return
	case "POST":
		username := r.FormValue("username")
		password := r.FormValue("password")
		// TODO: use hashed passwords. This is an example.
		ok, err := CheckCredentials(db, username, password)
		if err != nil {
			log.Printf("auth error: %v", err)
			http.Error(w, "Internal error", 500)
			return
		}
		if !ok {
			http.Redirect(w, r, "/login?failed=1", http.StatusFound)
			return
		}
		sess, _ := store.Get(r, "session-name")
		sess.Values["user"] = username
		sess.Values["authed_at"] = time.Now()
		_ = sess.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	sess, _ := store.Get(r, "session-name")
	delete(sess.Values, "user")
	_ = sess.Save(r, w)
	http.Redirect(w, r, "/", http.StatusFound)
}

// simple credential check (replace with hashed lookup)
func CheckCredentials(db *sql.DB, username, password string) (bool, error) {
	var storedHash string
	// example query: replace table/column names according to your schema
	err := db.QueryRow("SELECT password_hash FROM users WHERE username=$1", username).Scan(&storedHash)
	if err == sql.ErrNoRows {
		return false, nil
	} else if err != nil {
		return false, err
	}
	// TODO: compare hashed password (bcrypt)
	// For demo only: compare plaintext (DO NOT DO IN PROD)
	return password == storedHash, nil
}
