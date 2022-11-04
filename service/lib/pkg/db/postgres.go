package db

import (
	"database/sql"
	"fmt"

	"new-service/config"
	"new-service/lib/pkg/logger"

	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg *config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresDB.Host,
		cfg.PostgresDB.Port,
		cfg.PostgresDB.User,
		cfg.PostgresDB.Password,
		cfg.PostgresDB.Database,
	)

	logger.GetLogger().Info(fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s sslmode=disable",
		cfg.PostgresDB.Host,
		cfg.PostgresDB.Port,
		cfg.PostgresDB.User,
		cfg.PostgresDB.Database,
	))

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	return db, nil
}
