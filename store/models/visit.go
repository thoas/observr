package models

import (
	"database/sql"
	"time"
)

type Visit struct {
	Id          string         `db:"id"`
	Host        string         `db:"host"`
	Path        string         `db:"path"`
	RemoteAddr  string         `db:"remote_addr"`
	Method      string         `db:"method"`
	UserAgent   sql.NullString `db:"user_agent"`
	StatusCode  int            `db:"status_code"`
	Protocol    string         `db:"protocol"`
	Data        sql.NullString `db:"data"`
	Headers     sql.NullString `db:"headers"`
	Cookies     sql.NullString `db:"cookies"`
	Referer     sql.NullString `db:"referer"`
	QueryString sql.NullString `db:"query_string"`
	ProjectId   string         `db:"project_id"`
	Project     *Project       `db:"observr_project"`
	CreatedAt   time.Time      `db:"created_at"`
}

func (Visit) TableName() string {
	return "observr_visit"
}
