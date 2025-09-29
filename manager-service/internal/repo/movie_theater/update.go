package movietheater_repo

import (
	"sale-tickets/manager-service/internal/view"
	"time"
)

func (r *movietheaterRepo) Update(req view.UpdateMovieTheaterReq) error {
	queryStr := `
		UPDATE movie_theaters
		SET
			name = ?,
			address = ?,
			updated_at = ?
		WHERE uuid = ? AND deleted_at IS NULL
	`

	err := r.db.Exec(
		queryStr,
		req.Name,
		req.Address,
		time.Now(),
		req.Uuid,
	).Error
	if err != nil {
		return err
	}

	return nil
}
