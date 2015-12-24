package models

import (
	"github.com/gocql/gocql"
	"time"
)

type Visit struct {
	Id          gocql.UUID
	Host        string
	Path        string
	RemoteAddr  string
	Method      string
	UserAgent   string
	StatusCode  int
	Protocol    string
	Data        string
	Headers     string
	Cookies     string
	Referer     string
	QueryString string
	ProjectId   gocql.UUID
	CreatedAt   time.Time
}

func (m Visit) TableName() string {
	return "visits"
}

func (m Visit) PartitionKeys() []string {
	return []string{"Id"}
}
