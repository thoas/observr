package models

import (
	"github.com/gocql/gocql"
	"time"
)

type Tag struct {
	Id        gocql.UUID
	Key       string
	Value     string
	Data      string
	CreatedAt time.Time
	FirstSeen time.Time
	LastSeen  time.Time
	ProjectId gocql.UUID
	SeenCount time.Time
}

func (m Tag) TableName() string {
	return "tags"
}

func (m Tag) PartitionKeys() []string {
	return []string{"Id"}
}

type VisitTag struct {
	VisitId   gocql.UUID
	TagId     gocql.UUID
	CreatedAt time.Time
}

func (m VisitTag) TableName() string {
	return "visit_tags"
}

func (m VisitTag) PartitionKeys() []string {
	return []string{"Id"}
}
