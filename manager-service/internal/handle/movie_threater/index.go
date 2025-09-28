package moviethreater_controller

import (
	manager_api "sale-tickets/manager-service/gen"
	movietheater_service "sale-tickets/manager-service/internal/service/movie_theater"
)

type movieTheaterController struct {
	movieTheaterService movietheater_service.MovietheaterService
	manager_api.UnimplementedMovieTheaterServer
}

func NewHandle() manager_api.MovieTheaterServer {
	return &movieTheaterController{
		movieTheaterService: movietheater_service.NewMovietheaterService(),
	}
}
