package moviethreater_controller

import (
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/connection"

	"gorm.io/gorm"
)

type movieTheaterController struct {
	db *gorm.DB
	manager_api.UnimplementedMovieTheaterServer
}

func NewHandle() manager_api.MovieTheaterServer {
	return &movieTheaterController{
		db: connection.ConfigInfo.Database.GetConection(),
	}
}
