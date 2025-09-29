package movietheater_service

import (
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (s *movieTheaterService) Update(req view.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error) {
	err := s.movieTheaterRepo.Update(req)
	if err != nil {
		return nil, err
	}

	result := &manager_api.UpdateMovieTheaterRes{}
	return result, nil
}
