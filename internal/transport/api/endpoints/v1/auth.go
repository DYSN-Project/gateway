package v1

import (
	v1 "github.com/DYSN-Project/gateway/internal/transport/api/controller/v1"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
)

func RegisterAuthEndpoints(router *gin.Engine, grpcClient *grpc.Auth) {
	controller := v1.NewAuthController(grpcClient)
	grp := router.Group("api/v1/auth")

	grp.POST("/register", controller.ConfirmRegister)
	grp.POST("/confirm/", controller.ConfirmRegister)
	grp.POST("/login", controller.Login)
	grp.POST("/refresh-tokens", controller.RefreshTokens)
	grp.POST("/recovery-password", controller.RecoveryPassword)
	grp.POST("/confirm-recovery-password", controller.RecoveryPasswordConfirm)
	grp.POST("/change-password", controller.ChangePassword)
}
