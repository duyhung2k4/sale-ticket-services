package movietheater_repo

import (
	"sale-tickets/manager-service/internal/model"
	"sale-tickets/manager-service/internal/view"
)

func (r *movietheaterRepo) Detail(req view.DetailMovieTheaterReq) (*model.MovieTheater, error) {
	movieTheater := &model.MovieTheater{}

	queryStr := `
		SELECT
			uuid,
			name,
			address,
			creater_id,
			created_at,
			updated_at
		FROM movie_theaters
		WHERE 
			uuid = ?
				AND 
			deleted_at IS NULL
		LIMIT 1
	`

	err := r.db.Raw(
		queryStr,
		req.Uuid,
	).Scan(&movieTheater).Error
	if err != nil {
		return nil, err
	}

	return movieTheater, nil
}
