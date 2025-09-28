package view

import (
	"errors"
	manager_api "sale-tickets/manager-service/gen"
)

type CreateMovieTheaterReq struct {
	*manager_api.CreateMovieTheaterReq
	CreaterId string
}

func (v *CreateMovieTheaterReq) Validate() error {
	if v.Name == "" {
		return errors.New("invalid name")
	}

	if v.Address == "" {
		return errors.New("address name")
	}

	return nil
}

type GetListMovieTheaterReq struct {
	*manager_api.GetListMovieTheaterReq
}

func (v *GetListMovieTheaterReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.Filter == nil {
		return errors.New("filter data nil")
	}
	if v.Limit < 0 {
		return errors.New("invalid limit")
	}
	if v.Offset < 0 {
		return errors.New("invalid offset")
	}

	return nil
}
