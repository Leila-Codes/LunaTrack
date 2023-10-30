package main

import (
	"Luna_Track/db"
	"Luna_Track/log"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

const serverAddr = ":8080"

var (
	logger   = log.GetLogger()
	database db.IDatabase
)

func main() {
	logger.Debugf("Application starting...")

	database = db.ConnectDatabase(db.Postgres, db.DatabaseConfig{
		Host: "localhost",
		User: "postgres",
		Pass: "postgres",
		Name: "luna_track",
		Port: 5432,
	})

	router := mux.NewRouter()

	router.HandleFunc("/projects", ProjectListing)
	router.HandleFunc("/", index)

	router.PathPrefix("/").Handler(http.FileServer(http.Dir("www")))

	router.Use(log.LoggingMiddleware)

	logger.WithField("Addr", serverAddr).Info("Application Started")

	http.ListenAndServe(serverAddr, router)
}

// index page
func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("www/layout.html", "www/index.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		logger.Fatal(err)
	}
}
