package handler

import (
	"context"
	"errors"
	"new-service/domain/new/model"
	newpb "new-service/pb"
)

func (h *Handler) GetListData(ctx context.Context, msg *newpb.GetListDataRequest) (*newpb.GetListDataReturn, error) {
	select {
	case <-ctx.Done():
		return nil, errors.New("Abort")
	default:
	}

	res, err := h.usecase.GetListData(ctx, model.DataRequest{
		Search:     msg.Search,
		Restaurant: msg.Restaurant,
		LowPrice:   msg.LowPrice,
		HighPrice:  msg.HighPrice,
	})
	if err != nil {
		return nil, err
	}

	ret := newpb.GetListDataReturn{}
	for _, v := range res {
		ret.Lists = append(ret.Lists, &newpb.GetDataReturn{
			Id:   v.Id,
			Name: v.Name,
			Restaurant: &newpb.Restaurant{
				Id:      v.RestaurantId,
				Name:    v.RestaurantName,
				Address: v.RestaurantAddress,
			},
			Price: v.Price,
		})
	}

	return &ret, nil
}
