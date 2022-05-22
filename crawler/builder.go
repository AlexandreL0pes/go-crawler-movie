package crawler

import (
	"go-crawler-movie/crawler/config"
	"go-crawler-movie/domain/entities"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

func Build(h *colly.HTMLElement) entities.Movie {
	movie := entities.Movie{
		Id:         "",
		Title:      "",
		Year:       "",
		Rating:     0.0,
		Synopsis:   "",
		Categories: []string{},
		Stars:      make([]string, 0),
		Director:   "",
	}

	moviePath := getElementPath(h)

	movie.Id = h.ChildAttr(moviePath.Id, "data-tconst")
	movie.Title = h.ChildText(moviePath.Title)
	movie.Year = h.ChildText(moviePath.Year)
	movie.Synopsis = h.ChildText(moviePath.Synopsis)
	movie.Categories = strings.Split(h.ChildText(moviePath.Categories), ", ")
	movie.Director = h.ChildText(moviePath.Director)
	h.ForEach(moviePath.Stars, func(i int, h *colly.HTMLElement) {
		movie.Stars = append(movie.Stars, h.Text)
	})

	rating := h.ChildText(moviePath.Rating)

	if rating != "" {
		converted_rating, err := strconv.ParseFloat(rating, 32)

		if err == nil {
			movie.Rating = converted_rating
		}
	}

	return movie
}

func getElementPath(h *colly.HTMLElement) config.MovieInfoPath {
	hasRatingSection := func() bool {
		rating := h.ChildText(config.WithRatingPath.Rating)
		return rating != ""
	}

	if hasRatingSection() {
		return config.WithRatingPath
	}

	return config.WithoutRating
}
