package new

import (
	"context"
	"new-service/domain/new/model"
)

type NewRepoInterface interface {
	GetListData(ctx context.Context, in model.DataRequest) ([]model.DataResponse, error)
}
