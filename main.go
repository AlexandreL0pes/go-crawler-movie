package main

import (
	"fmt"
	"go-crawler-movie/crawler"
	"go-crawler-movie/database/dynamo"
	"go-crawler-movie/database/dynamo/repository"
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

	// setup dynamo db connection and movies repository
	dynamo, _ := dynamo.NewDynamoDB()
	repository := repository.Initialize(dynamo)

	// setup crawler and start crawling
	c := crawler.Initialize(repository)
	c.Execute()

	elapsedTime(start)
}

func elapsedTime(start time.Time) {
	elapsed := time.Since(start)
	fmt.Printf("\n\nTook around %s \n", elapsed)
}
