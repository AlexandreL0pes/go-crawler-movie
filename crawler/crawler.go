package crawler

import (
	"fmt"
	"go-crawler-movie/database/sqlite/repository"
	"time"

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/debug"
	"github.com/gocolly/colly/extensions"
)

const MAX_DEPTH = 2
const ASYNC_ENABLED = true

const totalIterations = 1

var currentInteration = 0

var url string = "https://www.imdb.com/search/title/?title_type=feature,tv_movie&count=250"

type Crawler struct {
	Collector     *colly.Collector
	Repository    *repository.MoviesRepository
	LogRepository *repository.ExecutionLogRepository
}

var crawler Crawler

func Initialize(mr *repository.MoviesRepository, lr *repository.ExecutionLogRepository) Crawler {
	crawler.Collector = colly.NewCollector(
		colly.MaxDepth(MAX_DEPTH),
		colly.Async(ASYNC_ENABLED),
		colly.Debugger(&debug.LogDebugger{}),
	)

	crawler.Repository = mr
	crawler.LogRepository = lr

	extensions.RandomUserAgent(crawler.Collector)
	extensions.Referer(crawler.Collector)

	crawler.Collector.Limit(&colly.LimitRule{
		DomainGlob:  "*",
		Parallelism: 4,
		RandomDelay: 2 * time.Second,
	})

	crawler.logOnRequest()
	crawler.setLanguageHeader()
	crawler.handleError()
	crawler.getMovies()
	crawler.navigate()

	return crawler
}

func (c Crawler) logOnRequest() {
	c.Collector.OnRequest(func(r *colly.Request) {
		fmt.Printf("\n\n")
		fmt.Printf("Crawling a new page...")
		fmt.Println(r.URL)
		fmt.Printf(">> Interation: %d from %d", currentInteration, totalIterations)
		fmt.Printf("\n>> Movies collected: %d\n", currentInteration*250)
		fmt.Printf("\n\n")
	})
}

func (c Crawler) setLanguageHeader() {
	c.Collector.OnRequest(func(r *colly.Request) {
		r.Headers.Set("Accept-Language", "en-US")
	})
}

func (c Crawler) handleError() {
	c.Collector.OnError(func(r *colly.Response, err error) {
		retryRequest(r.Request, 5)
		fmt.Println("\nRequest URL:", r.Request.URL, "failed with response:", r.StatusCode, "\nError:", err)
		fmt.Println("body: ", r.Body)
	})
}

func (c Crawler) getMovies() {
	c.Collector.OnHTML(".lister-item", func(h *colly.HTMLElement) {
		movie := Build(h)
		c.Repository.Insert(movie)
	})
}

func (c Crawler) Execute() {
	last_processed_url := c.LogRepository.GetLastExecution()
	fmt.Println("\n\nlast_processed_url: ", last_processed_url)
	if last_processed_url != "" {
		url = last_processed_url
	}
	c.Collector.Visit(url)
	c.Collector.Wait()
}

func (c Crawler) navigate() {
	c.Collector.OnHTML("a.lister-page-next", func(e *colly.HTMLElement) {
		if true {
			nextPage := e.Request.AbsoluteURL(e.Attr("href"))
			currentInteration++
			c.LogRepository.RegisterLastExecution(nextPage)
			c.Collector.Visit(nextPage)
		}
	})
}

func retryRequest(r *colly.Request, maxRetries int) int {
	retriesLeft := maxRetries
	if x, ok := r.Ctx.GetAny("retriesLeft").(int); ok {
		retriesLeft = x
	}
	if retriesLeft > 0 {
		r.Ctx.Put("retriesLeft", retriesLeft-1)
		r.Retry()
	}
	return retriesLeft
}
