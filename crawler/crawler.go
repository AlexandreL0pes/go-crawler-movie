package crawler

import (
	"fmt"
	"go-crawler-movie/domain/entity"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
)

const MAX_DEPTH = 2
const ASYNC_ENABLED = true

const totalIterations = 5

var currentInteration = 0

type Crawler struct {
	collector *colly.Collector
	Movies    []entity.Movie
}

var crawler Crawler

func Initialize() *colly.Collector {
	crawler.collector = colly.NewCollector(
		colly.MaxDepth(MAX_DEPTH),
		colly.Async(ASYNC_ENABLED),
	)

	extensions.RandomUserAgent(crawler.collector)
	extensions.Referer(crawler.collector)

	crawler.collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: 2 * time.Second,
	})

	crawler.logOnRequest()
	crawler.setLanguageHeader()
	crawler.handleError()

	return crawler.collector
}

func (c Crawler) logOnRequest() {
	c.collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("\n\n")
		fmt.Printf("Crawling a new page...")
		fmt.Println(r.URL)
		fmt.Printf(">> Interation: %d from %d", currentInteration, totalIterations)
		fmt.Printf("\n>> Movies collected: %d\n", currentInteration*250)
		fmt.Printf("\n\n")
	})
}

func (c Crawler) setLanguageHeader() {
	c.collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
	})
}

func (c Crawler) handleError() {
	c.collector.OnError(func(r *colly.Response, err error) {
		fmt.Println("\nRequest URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
		fmt.Println("body: ", r.Body)
	})
}

func (c Crawler) navigate() {
	c.collector.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		if currentInteration <= totalIterations {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			currentInteration++
			c.collector.Visit(nextPage)
		}
	})
}
