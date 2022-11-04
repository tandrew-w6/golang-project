package application

import (
	"context"
	"database/sql"
	"net"
	"new-service/config"
	"new-service/lib/pkg/db"
	"new-service/lib/pkg/logger"
	"strings"

	"github.com/urfave/cli"
	"google.golang.org/grpc"
)

type Application struct {
	Context    context.Context
	DbClient   *sql.DB
	GrpcServer *grpc.Server
}

func (app *Application) Run(cfg *config.Config) error {
	return func(app *Application) error {
		logger.GetLogger().Info("running grpc server...")
		lis, err := net.Listen("tcp", ":"+cfg.Application.ServerPort)
		if err != nil {
			logger.GetLogger().Error(err)
			return err
		}
		logger.GetLogger().Info(strings.ToUpper(cfg.Application.ServiceName) + " running on port " + cfg.Application.ServerPort)
		if err := app.GrpcServer.Serve(lis); err != nil {
			logger.GetLogger().Error(err)
			return err
		}
		app.GrpcServer.GracefulStop()
		return nil
	}(app)
}

func Setup(cfg *config.Config, c *cli.Context) (*Application, error) {
	app := new(Application)
	app.Context = context.Background()

	baseInit := []func(*Application) error{
		initDatabase(cfg),
		initGrpcServer(cfg),
	}

	for _, fn := range baseInit {
		if err := fn(app); err != nil {
			return app, err
		}
	}

	return app, nil
}

func initDatabase(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		db, err := db.NewPostgresDB(cfg)
		if err != nil {
			return err
		}

		app.DbClient = db
		logger.GetLogger().Info("init database done")

		return nil
	}
}

func initGrpcServer(cfg *config.Config) func(*Application) error {
	return func(app *Application) error {
		g := grpc.NewServer(
			grpc.ChainUnaryInterceptor(),
		)
		app.GrpcServer = g
		return nil
	}
}
