package usecase

import (
	"context"
	"new-service/domain/new/model"

	"github.com/sirupsen/logrus"
)

func (s *service) GetListData(ctx context.Context, req model.DataRequest) ([]model.DataResponse, error) {
	s.logger.WithFields(logrus.Fields{
		"Search":     req.Search,
		"Restaurant": req.Restaurant,
		"LowPrice":   req.LowPrice,
		"HighPrice":  req.HighPrice,
	}).Info("incoming request payload")

	results, err := s.repo.GetListData(ctx, req)
	if err != nil {
		return nil, err
	}

	return results, nil
}
