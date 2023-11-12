package models

type IssueType struct {
	ID         uint64
	ProjectKey string
	TypeName   string `sql:"type"`
}
