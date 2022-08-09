package v1

import (
	v1 "github.com/DYSN-Project/gateway/internal/transport/api/controller/v1"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterAuthEndpoints(router *gin.Engine, grpcClient *grpc.Auth, logger *logger.Logger) {
	controller := v1.NewAuthController(grpcClient, logger)

	grp := router.Group("api/v1/auth")

	grp.POST("/register", controller.ConfirmRegister)
	grp.GET("/confirm/:token", controller.ConfirmRegister)
	grp.POST("/login", controller.Login)
	grp.POST("/refresh-tokens", controller.RefreshTokens)
}
