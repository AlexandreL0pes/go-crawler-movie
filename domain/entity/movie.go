package entity

type Movie struct {
	Title      string   `json:"title"`
	Year       string   `json:"year"`
	Rating     float64  `json:"rating"`
	Synopsis   string   `json:"synopsis"`
	Categories []string `json:"categories"`
	Stars      []string `json:"stars"`
	Director   string   `json:"director"`
}
