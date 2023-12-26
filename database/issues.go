package database

import (
	"Luna_Track/database/models"
)

func SearchIssues(query models.IssueSearchQuery) ([]models.Issue, error) {
	qb := NewQueryBuilder("issues")
	qb.Select("project_key", "issue_id", "issue_type", "issue_status",
		"priority", "summary", "description",
		"created_at", "updated_at", "deleted_at")

	if len(query.ProjectKey) > 0 {
		qb.Where("project_key", query.ProjectKey)
	}

	if query.IssueTypeId != 0 {
		qb.Where("issue_type", query.IssueTypeId)
	}

	if len(query.StatusIds) > 0 {
		for _, status := range query.StatusIds {
			qb.Where("issue_status", status)
		}
	}

	if len(query.Term) > 0 {
		qb.Where("summary", query.Term)
	}
	// TODO: query.SummaryTerm
	// TODO: query.DescriptionTerm
	// TODO: multiple issue issue_types
	// TODO: allow 'LIKE' query searching

	rows, err := qb.Exec()
	if err != nil {
		return nil, err
	}

	results := make([]models.Issue, 0)
	for rows.Next() {
		issue := models.Issue{
			Type:     &models.IssueType{},
			Status:   &models.IssueStatus{},
			Priority: &models.IssuePriority{},
		}

		err = rows.Scan(
			&issue.ProjectKey,
			&issue.ID,
			&issue.Type.ID,
			&issue.Status.ID,
			&issue.Priority.ID,
			&issue.Summary,
			&issue.Description,
			&issue.CreatedAt,
			&issue.UpdatedAt,
			&issue.DeletedAt)

		if err != nil {
			return results, err
		}

		results = append(results, issue)
	}

	return results, nil
}

func ListIssues() ([]models.Issue, error) {
	rows, err := db.Query(`SELECT 
    project_key, issue_id, issue_type, issue_status, priority, summary, description, created_at, updated_at, deleted_at, created_by, U.username, U.display_name 
	FROM issues
	INNER JOIN users U on issues.created_by = U.user_id
	LIMIT 50`)
	if err != nil {
		return nil, err
	}

	results := make([]models.Issue, 0)
	for rows.Next() {
		issue := models.Issue{
			Type:      &models.IssueType{},
			Status:    &models.IssueStatus{},
			Priority:  &models.IssuePriority{},
			CreatedBy: &models.LunaUser{},
		}

		err = rows.Scan(
			&issue.ProjectKey,
			&issue.ID,
			&issue.Type.ID,
			&issue.Status.ID,
			&issue.Priority.ID,
			&issue.Summary,
			&issue.Description,
			&issue.CreatedAt,
			&issue.UpdatedAt,
			&issue.DeletedAt,
			&issue.CreatedBy.UserID,
			&issue.CreatedBy.Username,
			&issue.CreatedBy.DisplayName)

		if err != nil {
			return results, err
		}

		results = append(results, issue)
	}

	return results, nil
}

func GetIssueDetail(projectKey string, id uint64) (*models.Issue, error) {
	res := db.QueryRow(`SELECT
    I.project_key, I.issue_id, I.summary, I.description, I.created_at, I.updated_at, I.deleted_at,
    U.user_id, U.username, U.display_name,
    IType.issue_type_id, ITYpe.type,
    IStatus.issue_status_id, IStatus.status,
    IP.priority_id, IP.priority
FROM issues I
INNER JOIN users U ON I.created_by = U.user_id
INNER JOIN issue_priorities IP ON IP.project_key = I.project_key AND I.priority = IP.priority_id
INNER JOIN issue_types IType ON IType.project_key = I.project_key AND I.issue_type = IType.issue_type_id
INNER JOIN issue_statuses IStatus ON IStatus.project_key = I.project_key AND I.issue_status = IStatus.issue_status_id
WHERE I.project_key = $1 AND I.issue_id = $2;`, projectKey, id)

	if res.Err() != nil {
		return nil, res.Err()
	}

	issue := &models.Issue{
		Type:      &models.IssueType{},
		Status:    &models.IssueStatus{},
		Priority:  &models.IssuePriority{},
		CreatedBy: &models.LunaUser{},
		// TODO: Relationships: nil,
	}

	err := res.Scan(
		&issue.ProjectKey, &issue.ID, &issue.Summary, &issue.Description, &issue.CreatedAt, &issue.UpdatedAt, &issue.DeletedAt,
		&issue.CreatedBy.UserID, &issue.CreatedBy.Username, &issue.CreatedBy.DisplayName,
		&issue.Type.ID, &issue.Type.TypeName,
		&issue.Status.ID, &issue.Status.StatusText,
		&issue.Priority.ID, &issue.Priority.PriorityText,
	)

	return issue, err
}
