package v1

import (
	v1 "github.com/DYSN-Project/gateway/internal/transport/api/controller/v1"
	md "github.com/DYSN-Project/gateway/internal/transport/api/middleware"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterAreaEndpoints(router *gin.Engine,
	grpcClient *grpc.Area,
	grpcAuth *grpc.Auth,
	logger *logger.Logger) {
	controller := v1.NewAreaController(grpcClient, logger)
	middleware := md.NewAuthMdl(grpcAuth)

	grp := router.Group("api/v1/area")
	grp.GET("/", controller.List, middleware.Auth)
	grp.GET("/:id", controller.Show, middleware.Auth)
}
