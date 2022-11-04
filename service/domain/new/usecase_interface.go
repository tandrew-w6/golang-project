package new

import (
	"context"
	"new-service/domain/new/model"
)

type NewUsecaseInterface interface {
	GetListData(ctx context.Context, req model.DataRequest) ([]model.DataResponse, error)
}
