package v1

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
)

type AreaInterface interface {
	ListArea(c *gin.Context)
}

type AreaController struct {
	grpcClient *grpc.Area
}

func NewAreaController(grpcClient *grpc.Area) *AreaController {
	return &AreaController{
		grpcClient: grpcClient,
	}
}

func (a *AreaController) ListArea(c *gin.Context) {
	list, err := a.grpcClient.List()
	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, list)
}
