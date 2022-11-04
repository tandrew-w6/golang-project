package route

import (
	"new-service/application"
	"new-service/config"

	"new-service/domain/new/handler"
	"new-service/domain/new/repo"
	"new-service/domain/new/usecase"

	newpb "new-service/pb"
)

func SetupGrpcRouter(cfg *config.Config, app *application.Application) {
	postgresRepo := repo.NewPostgresRepo(app.DbClient)

	newUsecase := usecase.NewService(usecase.Dependencies{
		Repo: postgresRepo,
	})

	grpcHandler := handler.NewGrpcHandler(
		newUsecase,
	)

	newpb.RegisterNewServiceServer(app.GrpcServer, grpcHandler)
}
