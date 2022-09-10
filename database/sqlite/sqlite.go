package sqlite

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type SqliteDB struct {
	Client *gorm.DB
}

func NewSqliteDB() (*SqliteDB, error) {
	sqliteDB := &SqliteDB{}

	con, err := gorm.Open(sqlite.Open("movies.db"), &gorm.Config{})
	if err != nil {
		return sqliteDB, err
	}

	sqliteDB.Client = con

	return sqliteDB, nil
}

func (sdb *SqliteDB) Migrate(items interface{}) error {
	return sdb.Client.AutoMigrate(items)
}

func (sdb *SqliteDB) Insert(item interface{}) error {
	output := sdb.Client.Create(item)

	return output.Error
}

func (sdb *SqliteDB) BulkInsert(items interface{}) error {
	output := sdb.Client.CreateInBatches(items, 1000)

	return output.Error
}
