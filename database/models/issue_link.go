package models

type IssueLink struct {
	ProjectKey    string
	Relationship  *IssueLinkRelationship
	ParentIssueID uint64
	ParentIssue   *Issue
	ChildIssueID  uint64
	ChildIssue    *Issue
}
