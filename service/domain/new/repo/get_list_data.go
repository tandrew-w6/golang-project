package repo

import (
	"context"
	"fmt"

	"new-service/domain/new/model"
)

func (r *priceHunterRepo) GetListData(ctx context.Context, in model.DataRequest) ([]model.DataResponse, error) {
	datas := make([]model.DataResponse, 0)
	currBindParamCounter := 1
	paramQueries := []interface{}{}

	query := `
		SELECT md.id, md."name", md.restaurant_id, mr.name, mr.address, md.price  
		FROM m_dishes md 
		JOIN m_restaurants mr on mr.id = md.restaurant_id 
		WHERE md.deleted_at IS NULL AND mr.deleted_at IS NULL `

	if len(in.Search) > 0 {
		query += `AND (md.name ILIKE '%` + in.Search + `%' OR mr.name ILIKE '%` + in.Search + `%') `
	}

	if len(in.Restaurant) > 0 {
		query += fmt.Sprintf(`AND mr.name = $%d `, currBindParamCounter)
		paramQueries = append(paramQueries, in.Restaurant)
		currBindParamCounter += 1
	}

	if in.LowPrice != nil {
		query += fmt.Sprintf(`AND md.price >= $%d `, currBindParamCounter)
		paramQueries = append(paramQueries, in.LowPrice)
		currBindParamCounter += 1
	}

	if in.HighPrice != nil {
		query += fmt.Sprintf(`AND md.price <= $%d `, currBindParamCounter)
		paramQueries = append(paramQueries, in.HighPrice)
		currBindParamCounter += 1
	}

	rows, err := r.db.QueryContext(ctx, query, paramQueries...)
	if err != nil {
		r.log.Error(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var rowResult model.DataResponse
		err := rows.Scan(
			&rowResult.Id,
			&rowResult.Name,
			&rowResult.RestaurantId,
			&rowResult.RestaurantName,
			&rowResult.RestaurantAddress,
			&rowResult.Price,
		)
		if err != nil {
			r.log.Error(err)
			return nil, err
		}

		datas = append(datas, rowResult)
	}

	if len(datas) == 0 {
		return nil, nil
	}

	return datas, nil
}
