package models

type Model interface {
	TableName() string
	PartitionKeys() []string
}
