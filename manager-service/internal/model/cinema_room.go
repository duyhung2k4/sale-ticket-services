package model

type CinemaRoomModel struct {
	Uuid           string `json:"uuid"`
	MovieTheaterId string `json:"movie_theater_id"`
	Code           string `json:"code"`
}
