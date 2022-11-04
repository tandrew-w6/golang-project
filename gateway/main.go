package main

import (
	"context"
	"gateway-service/lib/pkg/logger"
	"gateway-service/route"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ctx := context.Background()
	log := logger.GetLogger()

	route.SetupRouterNew(ctx, log, router)
	router.Run("localhost:8080")
}
