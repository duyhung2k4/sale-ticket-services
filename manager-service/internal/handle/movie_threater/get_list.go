package moviethreater_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/view"
)

func (c *movieTheaterController) GetList(ctx context.Context, req *manager_api.GetListMovieTheaterReq) (*manager_api.GetListMovieTheaterRes, error) {
	reqData := view.GetListMovieTheaterReq{GetListMovieTheaterReq: req}
	if err := reqData.Validate(); err != nil {
		return nil, err
	}

	return &manager_api.GetListMovieTheaterRes{}, nil
}
