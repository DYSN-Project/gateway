package v1

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/model/form"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AreaInterface interface {
	List(c *gin.Context)
	Show(c *gin.Context)
}

type AreaController struct {
	grpcClient *grpc.Area
	logger     *logger.Logger
}

func NewAreaController(grpcClient *grpc.Area, logger *logger.Logger) *AreaController {
	return &AreaController{
		grpcClient: grpcClient,
		logger:     logger,
	}
}

func (a *AreaController) List(c *gin.Context) {
	rqt := &form.AreaList{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	list, err := a.grpcClient.List(rqt.Limit, rqt.Offset)
	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, list)
}

func (a *AreaController) Show(c *gin.Context) {
	id := c.Param("id")

	list, err := a.grpcClient.Show(id)
	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, list)
}
