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
		Update(req view.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error)
		GetList(req view.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error)
		Detail(req view.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error)
	}
)

func NewMovietheaterService() MovietheaterService {
	return &movieTheaterService{
		movieTheaterRepo: movietheater_repo.NewMovietheaterRepo(),
	}
}
