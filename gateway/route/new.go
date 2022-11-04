package route

import (
	"context"
	"net/http"

	newHandler "gateway-service/domain/new/handler/http_handler"
	newpb "gateway-service/pb"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func SetupRouterNew(ctx context.Context, log *logrus.Logger, r *gin.Engine) {
	grpcAddress := "localhost:10001"
	grpcConn, err := grpc.Dial(
		grpcAddress,
		grpc.WithInsecure(),
		grpc.WithChainUnaryInterceptor())
	if err != nil {
		log.Errorf("did not connect: %s", err)
	} else {
		log.Info("connected on " + grpcAddress)
	}
	grpcService := newpb.NewNewServiceClient(grpcConn)

	httpClient := &http.Client{}
	httpHandler := newHandler.NewHandler(ctx, log, grpcService, httpClient)

	r.GET("/datas", httpHandler.GetListData)
}
