package models

type IssueStatus struct {
	ID         uint64
	ProjectKey string
	StatusText string `sql:"status"`
}
