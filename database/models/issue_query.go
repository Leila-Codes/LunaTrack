package models

type IssueSearchQuery struct {
	ProjectKey      string
	IssueTypeId     int
	StatusIds       []int
	SummaryTerm     string
	DescriptionTerm string
	Term            string
}
