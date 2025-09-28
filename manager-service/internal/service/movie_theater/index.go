package movietheater_service

import (
	manager_api "sale-tickets/manager-service/gen"
	movietheater_repo "sale-tickets/manager-service/internal/repo/movie_theater"
	"sale-tickets/manager-service/internal/view"
)

type (
	movieTheaterService struct {
		movieTheaterRepo movietheater_repo.MovietheaterRepo
	}
	MovietheaterService interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
		GetList(req view.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error)
	}
)

func NewMovietheaterService() MovietheaterService {
	return &movieTheaterService{
		movieTheaterRepo: movietheater_repo.NewMovietheaterRepo(),
	}
}
