package db

import "Luna_Track/db/models"

type IDatabase interface {
	GetProjects() []models.ProjectListing
	GetProject(key string) models.ProjectDetails
	CreateProject(models.ProjectListing) models.ProjectListing
	UpdateProject(string, models.ProjectListing) bool
	DeleteProject(string) bool
}
