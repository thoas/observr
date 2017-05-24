package models

import (
	"database/sql"
	"time"

	"github.com/lib/pq"
)

type Tag struct {
	Id          string         `db:"id"`
	Key         string         `db:"key"`
	Value       string         `db:"value"`
	Data        sql.NullString `db:"data"`
	CreatedAt   time.Time      `db:"created_at"`
	FirstSeenAt pq.NullTime    `db:"first_seen_at"`
	LastSeenAt  pq.NullTime    `db:"last_seen_at"`
	ProjectId   string         `db:"project_id"`
	Project     *Project       `db:"observr_project"`
	SeenCount   time.Time      `db:"seen_count"`
}

func (Tag) TableName() string {
	return "observr_tag"
}

type VisitTag struct {
	Id        string `db:"id"`
	VisitId   string
	TagId     string
	CreatedAt time.Time
}

func (VisitTag) TableName() string {
	return "observr_visit_tag"
}

type GroupTag struct {
	Id        string `db:"id"`
	SrcTagId  string
	SrcTag    *Tag `db:"src_tag"`
	DstTagId  string
	DstTag    *Tag `db:"dst_tag"`
	CreatedAt time.Time
}

func (GroupTag) TableName() string {
	return "observr_group_tag"
}
