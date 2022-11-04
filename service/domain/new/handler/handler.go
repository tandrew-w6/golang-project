package handler

import (
	"new-service/domain/new"
	"new-service/lib/pkg/logger"

	"github.com/sirupsen/logrus"
)

type Handler struct {
	log     *logrus.Logger
	usecase new.NewUsecaseInterface
}

func NewGrpcHandler(usecase new.NewUsecaseInterface) *Handler {
	return &Handler{
		log:     logger.GetLogger(),
		usecase: usecase,
	}
}
