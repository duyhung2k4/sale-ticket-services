package moviethreater_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (c *movieTheaterController) Update(ctx context.Context, req *manager_api.UpdateMovieTheaterReq) (*manager_api.UpdateMovieTheaterRes, error) {
	reqData := view.UpdateMovieTheaterReq{UpdateMovieTheaterReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	result, err := c.movieTheaterService.Update(reqData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
