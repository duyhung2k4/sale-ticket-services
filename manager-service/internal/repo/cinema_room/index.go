package cinemaroom_repo

import (
	"sale-tickets/manager-service/internal/common/connection"
	"sale-tickets/manager-service/internal/view"

	"gorm.io/gorm"
)

type (
	cinemaRoomRepo struct {
		db *gorm.DB
	}
	CinemaRoomRepo interface {
		Create(req *view.CreateCinemaRoomReq) (uuid string, err error)
	}
)

func NewCinemaRoomRepo() CinemaRoomRepo {
	return &cinemaRoomRepo{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
