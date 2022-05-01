package main

import (
	"encoding/json"
	"fmt"
	"go-crawler-movie/config"
	"go-crawler-movie/crawler"
	"go-crawler-movie/domain/entity"
	"io/ioutil"
	"strconv"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

// var url string = "https://www.imdb.com/title/tt1877830/?ref_=watch_fanfav_tt_t_1"

// var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250"
// 620 interations - 155000 movies
// var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250&after=WzEyMzA0MCwidHQxMDU4MDE4OCIsNzc1MDFd&ref_=adv_nxt"
// 27000
// 2500

var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250&after=WzE0NzQ2MCwidHQwMDk5NTQ0Iiw5MTAwMV0%3D&ref_=adv_nxt"

const totalIterations = 2

var currentInteration = 0

func main() {
	fmt.Println("> executing crawler")
	scrape(url)
}

func scrape(url string) {
	start := time.Now()

	movies := []entity.Movie{}

	c := crawler.Initialize()

	c.OnHTML(".lister-item", func(h *colly.HTMLElement) {
		title := h.ChildText(config.Paths.Title)
		year := h.ChildText(config.Paths.Year)
		rating_s := h.ChildText(config.Paths.Rating)

		rating := 0.0
		synopsis := ""
		categories := []string{}
		director := ""
		stars := make([]string, 0)
		if rating_s != "" {
			a, err := strconv.ParseFloat(rating_s, 32)
			rating = a
			if err != nil {
				rating = 0.0
			}

			synopsis = h.ChildText(config.Paths.Synopsis)
			categories = strings.Split(h.ChildText(config.Paths.Categories), ", ")
			director = h.ChildText(config.Paths.Director)

			h.ForEach(config.Paths.Stars, func(i int, h *colly.HTMLElement) {
				stars = append(stars, h.Text)
			})
		} else {
			rating = 0.0

			synopsis = h.ChildText("div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(3)")
			categories = strings.Split(h.ChildText(config.Paths.Categories), ", ")
			director = h.ChildText("div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(4) > a:nth-child(1)")

			h.ForEach("div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(4) > span ~ a", func(i int, h *colly.HTMLElement) {
				stars = append(stars, h.Text)
			})

		}

		movies = append(movies, entity.Movie{title, year, rating, synopsis, categories, stars, director})

	})

	c.Visit(url)
	c.Wait()
	fmt.Printf("\n\nTook around %s \n", elapsedTime(start))
	storeMovies(movies)
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
