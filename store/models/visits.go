package models

import (
	"time"
)

type Visit struct {
	Id          string
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
	ProjectId   string
	CreatedAt   time.Time
}

func (m Visit) TableName() string {
	return "observr_visit"
}
