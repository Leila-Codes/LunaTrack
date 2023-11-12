package models

type Project struct {
	Key         string
	Name        string
	Description *string
	Owner       *LunaUser
}
