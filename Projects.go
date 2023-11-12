package main

import (
	"Luna_Track/database"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

func ProjectsApi(r *mux.Router) {
	r.Path("/{projectKey}").Methods(http.MethodGet).HandlerFunc(GetProject)
	r.Methods(http.MethodGet).HandlerFunc(ListProjects)
}

func ListProjects(w http.ResponseWriter, r *http.Request) {
	projectList, err := database.ListProjects()
	if err != nil {
		HttpError(w, http.StatusInternalServerError, err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/projects/listing.gohtml"))
	err = tmpl.Execute(w, projectList)
	if err != nil {
		logger.Fatal(err)
	}
}

func GetProject(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	projectDetail, err := database.GetProject(vars["projectKey"])
	if err != nil {
		HttpError(w, http.StatusInternalServerError, err)
		return
	}

	tmpl := template.Must(template.ParseFiles("templates/projects/details.gohtml"))
	err = tmpl.Execute(w, projectDetail)
	if err != nil {
		HttpError(w, http.StatusInternalServerError, err)
	}
}
