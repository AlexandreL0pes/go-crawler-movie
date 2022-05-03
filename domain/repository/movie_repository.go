package repository

import (
	"go-crawler-movie/domain/entity"
)

type MoviesRepository struct {
	Movies []entity.Movie
}

func Initialize() MoviesRepository {
	return MoviesRepository{}
}

func (s *MoviesRepository) Save(movie entity.Movie) {
	s.Movies = append(s.Movies, movie)
}
