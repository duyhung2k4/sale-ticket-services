package moviethreater_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (c *movieTheaterController) Create(ctx context.Context, req *manager_api.CreateMovieTheaterReq) (*manager_api.CreateMovieTheaterRes, error) {
	reqData := view.CreateMovieTheaterReq{CreateMovieTheaterReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	id, err := c.movieTheaterService.Create(reqData)
	if err != nil {
		return nil, err
	}

	response := &manager_api.CreateMovieTheaterRes{
		Uuid:    id,
		Name:    reqData.Name,
		Address: reqData.Address,
	}
	return response, nil
}
