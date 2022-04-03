package main

import (
	"encoding/json"
	"fmt"
	"go-crawler-movie/config"
	"io/ioutil"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// var url string = "https://www.imdb.com/title/tt1877830/?ref_=watch_fanfav_tt_t_1"
var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250"

const totalIterations = 10

var currentInteration = 0

type Movie struct {
	Title      string   `json:"title"`
	Year       string   `json:"year"`
	Rating     float64  `json:"rating"`
	Synopsis   string   `json:"synopsis"`
	Categories []string `json:"categories"`
	Stars      []string `json:"stars"`
	Director   string   `json:"director"`
}

func main() {
	fmt.Println("> executing crawler")
	scrape(url)
}

func scrape(url string) {
	movies := []Movie{}
	c := colly.NewCollector()

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("> Visiting", r.URL)
		fmt.Printf(">> Interation: %d from %d", currentInteration, totalIterations)
		fmt.Printf("\n>> Movies collected: %d", currentInteration*250)
	})

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

		movies = append(movies, Movie{title, year, rating, synopsis, categories, stars, director})

	})

	c.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		if currentInteration <= totalIterations {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			currentInteration++
			c.Visit(nextPage)
		}
	})

	c.Visit(url)

	storeMovies(movies)
}

func storeMovies(m []Movie) {
	file, _ := json.MarshalIndent(m, "", " ")
	_ = ioutil.WriteFile("movies.json", file, 0644)
}
