package movietheater_repo

import (
	"sale-tickets/manager-service/internal/view"
)

func (r *movietheaterRepo) Update(req view.UpdateMovieTheaterReq) error {
	queryStr := `
		UPDATE movie_theaters
		SET
			name = ?,
			address = ?
		WHERE uuid = ? AND deleted_at IS NULL
	`

	err := r.db.Exec(
		queryStr,
		req.Name,
		req.Address,
		req.Uuid,
	).Error
	if err != nil {
		return err
	}

	return nil
}
