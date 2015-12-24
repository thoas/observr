package store

import (
	"github.com/hailocab/gocassa"
	"github.com/thoas/observr/store/models"
)

type DataStore struct {
	KeySpace gocassa.KeySpace
	Tables   map[string]gocassa.Table
}

type Option struct {
	Name     string
	Ips      []string
	Username string
	Password string
}

func NewDataStore(option *Option) (*DataStore, error) {
	keySpace, err := gocassa.ConnectToKeySpace(option.Name, option.Ips, option.Username, option.Password)

	if err != nil {
		return nil, err
	}

	keySpace.DebugMode(true)

	tables := make(map[string]gocassa.Table)

	store := &DataStore{KeySpace: keySpace}

	for _, model := range store.Models() {
		table := store.KeySpace.Table(model.TableName(), model, gocassa.Keys{
			PartitionKeys: model.PartitionKeys(),
		}).WithOptions(gocassa.Options{
			TableName: model.TableName(),
		})

		tables[model.TableName()] = table
	}

	store.Tables = tables

	return store, nil
}

func (s *DataStore) RecreateTables() error {
	for _, table := range s.Tables {
		err := table.Recreate()

		if err != nil {
			return err
		}
	}

	return nil
}

func (s *DataStore) Models() []models.Model {
	return []models.Model{
		models.Visit{},
		models.VisitTag{},
		models.Tag{},
	}
}
