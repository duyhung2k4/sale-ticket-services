package model

type CinemaRoom struct {
	Base
	MovieTheaterId string `json:"movie_theater_id"`
	Code           string `json:"code"`
}
