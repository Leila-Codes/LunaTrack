package models

import "time"

type ProjectListing struct {
	Key, Title          string
	Description         *string
	Created, LastOpened time.Time
}
