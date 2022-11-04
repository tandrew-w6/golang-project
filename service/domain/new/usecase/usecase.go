package usecase

import (
	"new-service/lib/pkg/logger"

	"github.com/sirupsen/logrus"

	"new-service/domain/new"
)

type service struct {
	logger *logrus.Logger
	repo   new.NewRepoInterface
}

type Dependencies struct {
	Repo new.NewRepoInterface
}

func NewService(deps Dependencies) new.NewUsecaseInterface {
	return &service{
		logger: logger.GetLogger(),
		repo:   deps.Repo,
	}
}
