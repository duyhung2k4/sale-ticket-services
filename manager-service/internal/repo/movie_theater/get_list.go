package movietheater_repo

import (
	"fmt"
	"sale-tickets/manager-service/internal/common/utils"
	"sale-tickets/manager-service/internal/model"
	"sale-tickets/manager-service/internal/view"
)

func (r *movietheaterRepo) GetList(req view.GetListMovieTheaterReq) ([]model.MovieTheater, error) {
	results := []model.MovieTheater{}

	queryStr := fmt.Sprintf(`
		SELECT
			uuid,
			creater_id,
			name,
			address
		FROM movie_theaters
		WHERE
			%s
		LIMIT ?
		OFFSET ?
		`,
		utils.MergeConditionAND(
			utils.AddLikeClause("name"),
			utils.AddLikeClause("address"),
		),
	)

	err := r.db.Raw(
		queryStr,
		req.Filter.Name,
		req.Filter.Address,
		req.Limit,
		req.Offset,
	).Scan(&results).Error
	if err != nil {
		return nil, err
	}

	return results, nil
}
