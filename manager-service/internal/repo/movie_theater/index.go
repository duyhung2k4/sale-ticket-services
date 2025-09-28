package movietheater_repo

import (
	"sale-tickets/manager-service/internal/common/connection"
	"sale-tickets/manager-service/internal/view"

	"gorm.io/gorm"
)

type (
	movietheaterRepo struct {
		db *gorm.DB
	}
	MovietheaterRepo interface {
		Create(req view.CreateMovieTheaterReq) (id string, err error)
	}
)

func NewMovietheaterRepo() MovietheaterRepo {
	return &movietheaterRepo{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
