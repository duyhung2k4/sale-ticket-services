package movietheater_repo

import (
	"sale-tickets/manager-service/internal/model"
	"sale-tickets/manager-service/internal/view"

	"github.com/google/uuid"
)

func (r *movietheaterRepo) Create(req view.CreateMovieTheaterReq) (id string, err error) {
	uuidBuf, err := uuid.NewV7()
	if err != nil {
		return "", err
	}

	movie := model.MovieTheater{
		Base: model.Base{
			Uuid: uuidBuf.String(),
		},
		CreaterId: req.CreaterId,
		Name:      req.Name,
		Address:   req.Address,
	}
	err = r.db.Model(&model.MovieTheater{}).Create(&movie).Error
	if err != nil {
		return "", err
	}

	return movie.Uuid, nil
}
