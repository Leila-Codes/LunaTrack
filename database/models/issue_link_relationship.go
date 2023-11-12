package models

type IssueLinkRelationship struct {
	ProjectKey       string
	RelationshipID   uint64
	RelationshipText string `sql:"relationship"`
}
