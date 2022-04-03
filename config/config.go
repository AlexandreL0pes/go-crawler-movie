package config

type movieInfo struct {
	Title      string
	Year       string
	Rating     string
	Synopsis   string
	Categories string
	Stars      string
	Director   string
}

var Paths = movieInfo{
	Title:      "div.lister.list.detail.sub-list > div > div > div.lister-item-content > h3 > a",
	Year:       "div.lister.list.detail.sub-list > div > div > div.lister-item-content > h3 > span.lister-item-year.text-muted.unbold",
	Rating:     "div.lister.list.detail.sub-list > div > div > div.lister-item-content > div > div.inline-block.ratings-imdb-rating > strong",
	Synopsis:   "div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(4)",
	Categories: "div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(2) > span.genre",
	Stars:      "div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(5) > span ~ a",
	Director:   "div.lister.list.detail.sub-list > div > div > div.lister-item-content > p:nth-child(5) > a:nth-child(1)",
}
