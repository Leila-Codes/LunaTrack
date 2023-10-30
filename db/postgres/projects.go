package postgres

import "Luna_Track/db/models"

func (p *PostgesDB) GetProjects() []models.ProjectListing {
	rows, err := p.conn.Query("SELECT key, title, description FROM project_listing")
	if err != nil {
		panic(err)
	}

	projects := make([]models.ProjectListing, 0)

	for rows.Next() {
		var project models.ProjectListing
		err = rows.Scan(&project.Key, &project.Title, &project.Description)
		if err != nil {
			panic(err)
		}

		projects = append(projects, project)
	}

	return projects
}

func (p *PostgesDB) GetProject(key string) models.ProjectDetails {
	panic("not yet implemented")
}

func (p *PostgesDB) CreateProject(details models.ProjectListing) models.ProjectListing {
	panic("not yet implemented")
}

func (p *PostgesDB) UpdateProject(key string, details models.ProjectListing) bool {
	panic("not yet implemented")
}

func (p *PostgesDB) DeleteProject(key string) bool {
	panic("not yet implemented")
}
