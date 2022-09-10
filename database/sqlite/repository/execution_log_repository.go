package repository

import (
	"go-crawler-movie/database/sqlite"
	"time"

	"gorm.io/gorm"
)

type executionLog struct {
	FinishedAt       time.Time
	LastProcessedUrl string
	gorm.Model
}

type ExecutionLogRepository struct {
	db *sqlite.SqliteDB
}

func NewExecutionLogRepository(db *sqlite.SqliteDB) *ExecutionLogRepository {
	err := db.Migrate(&executionLog{})

	if err != nil {
		panic(err)
	}

	return &ExecutionLogRepository{
		db: db,
	}
}

func (elr *ExecutionLogRepository) RegisterLastExecution(url string) error {
	return elr.db.Insert(&executionLog{FinishedAt: time.Now(), LastProcessedUrl: url})
}

func (elr *ExecutionLogRepository) GetLastExecution() string {
	var lastExecutionLog executionLog
	elr.db.Client.Order("finished_at DESC").Limit(1).Find(&lastExecutionLog)

	return lastExecutionLog.LastProcessedUrl
}
