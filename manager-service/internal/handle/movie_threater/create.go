package moviethreater_controller

import (
	"context"
	manager_api "sale-tickets/manager-service/gen"
	"sale-tickets/manager-service/internal/model"

	"github.com/google/uuid"
)

func (s *movieTheaterController) Create(ctx context.Context, req *manager_api.CreateMovieTheaterReq) (*manager_api.CreateMovieTheaterRes, error) {
	uuidBuf, err := uuid.NewV7()
	if err != nil {
		return nil, err
	}

	payload := &model.MovieTheaterModel{
		Uuid:    uuidBuf.String(),
		Name:    req.Name,
		Address: req.Address,
	}
	err = s.db.Model(&model.MovieTheaterModel{}).Create(&payload).Error
	if err != nil {
		return nil, err
	}

	response := &manager_api.CreateMovieTheaterRes{
		Uuid:    payload.Uuid,
		Name:    payload.Name,
		Address: payload.Address,
	}
	return response, nil
}
