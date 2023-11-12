package models

type ProjectMember struct {
	ProjectKey string
	Project    *Project
	UserID     uint64
	User       *LunaUser
}
