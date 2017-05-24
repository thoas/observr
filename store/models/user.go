package models

import (
	"time"

	"github.com/lib/pq"
)

type User struct {
	Id        string      `db:"id"`
	Username  string      `db:"username"`
	Email     string      `db:"email"`
	Password  string      `db:"password"`
	CreatedAt time.Time   `db:"created_at"`
	UpdatedAt time.Time   `db:"updated_at"`
	DeletedAt pq.NullTime `db:"deleted_at"`
}

func (User) TableName() string {
	return "observr_user"
}
