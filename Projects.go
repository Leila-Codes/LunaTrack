package main

import (
	"html/template"
	"net/http"
	"time"
)

type ProjectItem struct {
	ProjectName        string
	ProjectDescription string
	ProjectDate        time.Time
}

var (
	projects = []ProjectItem{
		{ProjectName: "Test Project #1", ProjectDescription: "This is a test project for your amazement.", ProjectDate: time.Now()},
		{ProjectName: "Uber Project #2", ProjectDescription: "This is a huge project, yeet!!!", ProjectDate: time.Now().Add(-36 * time.Hour)},
		{ProjectName: "LEGACY project", ProjectDescription: "Something really, really old.", ProjectDate: time.Unix(0, 0)},
	}
)

func ProjectListing(w http.ResponseWriter, r *http.Request) {
	// TODO: Apply a filter to the projects (e.g. search)
	tmpl := template.Must(template.ParseFiles("templates/project-listing.html"))
	err := tmpl.Execute(w, projects)
	if err != nil {
		logger.Fatal(err)
	}
}
