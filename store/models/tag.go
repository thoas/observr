package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Tag struct {
	ID          string         `db:"id"`
	Key         string         `db:"key"`
	Value       string         `db:"value"`
	Data        sql.NullString `db:"data"`
	CreatedAt   time.Time      `db:"created_at"`
	FirstSeenAt pq.NullTime    `db:"first_seen_at"`
	LastSeenAt  pq.NullTime    `db:"last_seen_at"`
	ProjectID   string         `db:"project_id"`
	Project     *Project       `db:"observr_project"`
	SeenCount   time.Time      `db:"seen_count"`
}

func (Tag) TableName() string {
	return "observr_tag"
}

type VisitTag struct {
	Id        string `db:"id"`
	VisitID   string
	TagID     string
	CreatedAt time.Time
}

func (VisitTag) TableName() string {
	return "observr_visit_tag"
}

type GroupTag struct {
	Id        string `db:"id"`
	SrcTagID  string
	SrcTag    *Tag `db:"src_tag"`
	DstTagID  string
	DstTag    *Tag `db:"dst_tag"`
	CreatedAt time.Time
}

func (GroupTag) TableName() string {
	return "observr_group_tag"
}
