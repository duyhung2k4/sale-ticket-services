package view

import (
	"errors"

	manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"
)

type CreateCinemaRoomReq struct {
	*manager_api.CreateCinemaRoomReq
}

func (v *CreateCinemaRoomReq) Validate() error {
	if v == nil {
		return errors.New("request data nil")
	}
	if v.MovieTheaterId == "" {
		return errors.New("movie_theater_id invalid")
	}
	if v.Code == "" {
		return errors.New("code invalid")
	}
	return nil
}
