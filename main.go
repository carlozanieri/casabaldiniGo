package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
)

// global templates variable (parse all templates at startup)
var templates = template.Must(template.ParseGlob("templates/*.html"))

// session store (use env var SESSION_KEY)
var store *sessions.CookieStore

// DB handle
var db *sql.DB

func main() {
	// session key from env (must be 32 or 64 bytes ideally)
	sessionKey := os.Getenv("SESSION_KEY")
	if sessionKey == "" {
		log.Fatal("set SESSION_KEY env var")
	}
	store = sessions.NewCookieStore([]byte(sessionKey))
	store.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7, // 7 days
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
	}

	// connect DB (example via env variables)
	var err error
	db, err = OpenDBFromEnv()
	if err != nil {
		log.Fatalf("db open: %v", err)
	}
	defer db.Close()

	r := mux.NewRouter()

	// middleware chain
	r.Use(LoggingMiddleware)
	r.Use(RecoverMiddleware)
	r.Use(TimingMiddleware)

	// static
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// route groups (Flask-style): create subrouters
	public := r.PathPrefix("/").Subrouter()
	public.HandleFunc("/", HomeHandler).Methods("GET")
	public.HandleFunc("/about", AboutHandler).Methods("GET")
	public.HandleFunc("/contact", ContactHandler).Methods("GET")
	public.HandleFunc("/login", LoginHandler).Methods("GET", "POST")
	public.HandleFunc("/logout", LogoutHandler).Methods("POST")

	// protected group example
	api := r.PathPrefix("/admin").Subrouter()
	api.Use(AuthRequiredMiddleware) // add auth middleware to group
	api.HandleFunc("", AdminIndexHandler).Methods("GET")

	// custom 404
	r.NotFoundHandler = http.HandlerFunc(NotFoundHandler)

	srv := &http.Server{
		Handler:      r,
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 20 * time.Second,
		IdleTimeout:  60 * time.Second,
	}

	log.Println("Starting server on", srv.Addr)
	log.Fatal(srv.ListenAndServe())
}
