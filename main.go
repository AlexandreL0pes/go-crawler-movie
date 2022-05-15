package main

import (
	"encoding/json"
	"fmt"
	"go-crawler-movie/crawler"
	"go-crawler-movie/database/dynamo"
	"go-crawler-movie/domain/entity"
	"go-crawler-movie/domain/repository"
	"io/ioutil"
	"time"
)

// var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250"
// 620 interations - 155000 movies
// var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250&after=WzEyMzA0MCwidHQxMDU4MDE4OCIsNzc1MDFd&ref_=adv_nxt"
// 27000
// 2500

// var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250&after=WzE0NzQ2MCwidHQwMDk5NTQ0Iiw5MTAwMV0%3D&ref_=adv_nxt"

func main() {
	fmt.Println("> executing crawler")
	start := time.Now()
	dynamo, _ := dynamo.NewDynamoDB()
	repository := repository.Initialize(dynamo)

	c := crawler.Initialize(repository)
	c.Execute()

	storeMovies(repository.Movies)
	fmt.Printf("\n\nTook around %s \n", elapsedTime(start))
}

func storeMovies(m []entity.Movie) {
	file, _ := json.MarshalIndent(m, "", " ")
	now := time.Now()
	filename := fmt.Sprintf("movies teste - %s.json", now.Format(time.UnixDate))

	_ = ioutil.WriteFile(filename, file, 0644)
}

func elapsedTime(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}
