package repository

import (
	"encoding/json"
	"go-crawler-movie/domain/entities"

	"go-crawler-movie/database/sqlite"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

const TABLE_NAME = "movies"

type movie struct {
	ID      uuid.UUID
	Content string
	gorm.Model
}
type MoviesRepository struct {
	db *sqlite.SqliteDB
}

func Initialize(db *sqlite.SqliteDB) *MoviesRepository {
	err := db.Migrate(&movie{})
	if err != nil {
		panic(err)
	}
	return &MoviesRepository{
		db: db,
	}
}

func (s *MoviesRepository) Insert(m entities.Movie) error {
	mJson, err := json.Marshal(m)
	if err != nil {
		return nil
	}
	err = s.db.Insert(&movie{ID: uuid.New(), Content: string(mJson)})

	return err
}
