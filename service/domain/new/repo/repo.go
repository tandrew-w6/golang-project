package repo

import (
	"database/sql"
	"new-service/domain/new"
	"new-service/lib/pkg/logger"

	"github.com/sirupsen/logrus"
)

type priceHunterRepo struct {
	log *logrus.Logger
	db  *sql.DB
}

func NewPostgresRepo(db *sql.DB) new.NewRepoInterface {
	return &priceHunterRepo{
		log: logger.GetLogger(),
		db:  db,
	}
}
