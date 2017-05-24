package models

import (
	"time"
)

type Tag struct {
	Id          string
	Key         string
	Value       string
	Data        string
	CreatedAt   time.Time
	FirstSeenAt time.Time
	LastSeenAt  time.Time
	ProjectId   string
	SeenCount   time.Time
}

func (m Tag) TableName() string {
	return "observr_tag"
}

type VisitTag struct {
	VisitId   string
	TagId     string
	CreatedAt time.Time
}

func (m VisitTag) TableName() string {
	return "observr_visit_tag"
}
