package cinemaroom_controller

import manager_api "github.com/duyhung2k4/sale-tickets-golang-common/manager-api/proto"

type cinemaRoomController struct {
	manager_api.UnsafeCinemaRoomServiceServer
}

func NewCinemaRoomController() *cinemaRoomController {
	return &cinemaRoomController{}
}
