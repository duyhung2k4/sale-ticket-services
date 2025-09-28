package movietheater_repo

import (
	"sale-tickets/manager-service/internal/common/connection"
	"sale-tickets/manager-service/internal/model"
	"sale-tickets/manager-service/internal/view"

	"gorm.io/gorm"
)

type (
	movietheaterRepo struct {
		db *gorm.DB
	}
	MovietheaterRepo interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
		GetList(req view.GetListMovieTheaterReq) ([]model.MovieTheater, error)
		CountList(req view.GetListMovieTheaterReq) (int32, error)
	}
)

func NewMovietheaterRepo() MovietheaterRepo {
	return &movietheaterRepo{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
