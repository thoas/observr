package models

import (
	"time"

	"github.com/lib/pq"
)

type Project struct {
	Id        string      `db:"id"`
	Name      string      `db:"name"`
	URL       string      `db:"url"`
	ApiKey    string      `db:"api_key"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

func (Project) TableName() string {
	return "observr_project"
}
