package database

import (
	"Luna_Track/database/models"
	"fmt"
)

func ListProjects() (projects []models.Project, err error) {
	projects = make([]models.Project, 0)

	rows, err := db.Query(`SELECT key, name, description FROM projects`)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		proj := models.Project{}
		err = rows.Scan(&proj.Key, &proj.Name, &proj.Description)
		if err != nil {
			return nil, err
		}

		projects = append(projects, proj)
	}

	return projects, nil
}

func GetProject(key string) (project *models.Project, err error) {
	row := db.QueryRow(
		`SELECT p.key, p.name, p.description, u.username, u.display_name 
FROM projects p INNER JOIN users u ON p.owner_id = u.user_id 
WHERE p.key = $1`,
		key,
	)

	if row.Err() != nil {
		return nil, row.Err()
	}

	project = new(models.Project)
	user := new(models.LunaUser)

	err = row.Scan(
		&project.Key,
		&project.Name,
		&project.Description,
		&user.Username,
		&user.DisplayName,
	)
	if err != nil {
		return nil, err
	}

	project.Owner = user

	return project, nil
}

func UpdateProject(key string, data models.Project) error {
	res, err := db.Exec(`UPDATE projects SET name = $1, description = $2 WHERE key = $3`,
		data.Name, data.Description, key,
	)

	if err != nil {
		return err
	}

	if count, err := res.RowsAffected(); err != nil || count != 1 {
		return fmt.Errorf("failed to modify, no match or error %v", err)
	}

	return nil
}

func DeleteProject(key string) error {
	res, err := db.Exec(`DELETE FROM projects CASCADE WHERE key = $1`, key)
	if err != nil {
		return err
	}

	if count, err := res.RowsAffected(); err != nil || count != 1 {
		return fmt.Errorf("failed to modify, no match or error %v", err)
	}

	return nil
}
