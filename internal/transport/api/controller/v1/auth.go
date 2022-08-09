package v1

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/model/form"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Register(c gin.Context)
	ConfirmRegister(c gin.Context)
	Login(c gin.Context)
	RefreshTokens(c gin.Context)
}

type AuthController struct {
	grpcClient *grpc.Auth
	logger     *logger.Logger
}

func NewAuthController(grpcClient *grpc.Auth, logger *logger.Logger) *AuthController {
	return &AuthController{
		grpcClient: grpcClient,
		logger:     logger,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	rqt := &form.RegisterForm{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	token, err := a.grpcClient.Register(rqt.Email, rqt.Password)
	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, token)
}

func (a *AuthController) ConfirmRegister(c *gin.Context) {

	rqt := c.Param("token")
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	token, err := a.grpcClient.Register(rqt.GetTok)
	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, token)
}
