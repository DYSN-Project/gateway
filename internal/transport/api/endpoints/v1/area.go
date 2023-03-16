package v1

import (
	v1 "github.com/DYSN-Project/gateway/internal/transport/api/controller/v1"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
)

func RegisterAreaEndpoints(router *gin.Engine, grpcClient *grpc.Area) {
	controller := v1.NewAreaController(grpcClient)

	grp := router.Group("api/v1/area")
	grp.GET("/", controller.ListArea)
}
