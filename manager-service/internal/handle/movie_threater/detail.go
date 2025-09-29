package moviethreater_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (c *movieTheaterController) Detail(ctx context.Context, req *manager_api.DetailMovieTheaterReq) (*manager_api.DetailMovieTheaterRes, error) {
	reqData := view.DetailMovieTheaterReq{DetailMovieTheaterReq: req}

	result, err := c.movieTheaterService.Detail(reqData)
	if err != nil {
		return nil, err
	}

	return result, nil
}
