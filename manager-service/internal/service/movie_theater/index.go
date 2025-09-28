package movietheater_service

import (
	movietheater_repo "sale-tickets/manager-service/internal/repo/movie_theater"
	"sale-tickets/manager-service/internal/view"
)

type (
	movieTheaterService struct {
		movieTheaterRepo movietheater_repo.MovietheaterRepo
	}
	MovietheaterService interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
	}
)

func NewMovietheaterService() MovietheaterService {
	return &movieTheaterService{
		movieTheaterRepo: movietheater_repo.NewMovietheaterRepo(),
	}
}
