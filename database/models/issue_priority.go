package models

type IssuePriority struct {
	ID           uint64
	ProjectKey   string
	PriorityText string `sql:"priority"`
}
