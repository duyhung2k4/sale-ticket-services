package movietheater_service

import (
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (s *movieTheaterService) Detail(req view.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error) {
	movieTheater, err := s.movieTheaterRepo.Detail(req)
	if err != nil {
		return nil, err
	}

	result := &manager_api.DetailMovieTheaterRes{
		Uuid:      movieTheater.Uuid,
		Name:      movieTheater.Name,
		Address:   movieTheater.Address,
		CreaterId: movieTheater.CreaterId,
		CreatedAt: movieTheater.CreatedAt.String(),
		UpdatedAt: movieTheater.CreatedAt.String(),
	}

	return result, nil
}
