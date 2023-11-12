package models

import "time"

type Issue struct {
	ProjectKey    string
	Project       *Project
	Type          *IssueType
	Status        *IssueStatus
	Priority      *IssuePriority
	Summary       string
	Description   *string
	CreatedAt     time.Time
	UpdatedAt     *time.Time
	DeletedAt     *time.Time
	CreatedById   uint64
	CreatedBy     *LunaUser
	Relationships *[]IssueLinkRelationship
}
