package main

import (
	"fmt"
	"go-crawler-movie/crawler"
	"go-crawler-movie/database/sqlite/repository"
	"time"

	"go-crawler-movie/database/sqlite"
)

func main() {
	start := time.Now()

	sdb, err := sqlite.NewSqliteDB()
	if err != nil {
		panic(err)
	}
	moviesRepository := repository.Initialize(sdb)
	logsRepository := repository.NewExecutionLogRepository(sdb)
	// setup crawler and start crawling
	c := crawler.Initialize(moviesRepository, logsRepository)
	c.Execute()

	elapsedTime(start)
}

func elapsedTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("\n\nTook around %s \n", elapsed)
}
