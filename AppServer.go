package main

import (
	"Luna_Track/database"
	"Luna_Track/log"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

const serverAddr = ":8080"

var (
	logger = log.GetLogger()
)

func main() {
	logger.Debugf("Application starting...")

	err := database.Connect(
		"localhost",
		5432,
		"postgres",
		"postgres",
		"luna_track",
	)
	if err != nil {
		panic(err)
	}

	router := mux.NewRouter()

	projectsApi := router.PathPrefix("/projects").Subrouter()
	ProjectsApi(projectsApi)
	//router.Use(projectsApi)

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
