package movietheater_service

import "sale-tickets/manager-service/internal/view"

func (s *movieTheaterService) Create(req view.CreateMovieTheaterReq) (id string, err error) {
	return s.movieTheaterRepo.Create(req)
}
