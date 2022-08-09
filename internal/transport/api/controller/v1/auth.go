package v1

import (
	"errors"
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
	token := c.Param("token")

	user, err := a.grpcClient.ConfirmRegister(token)
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, user)
}

func (a *AuthController) Login(c *gin.Context) {
	rqt := &form.LoginForm{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	tokens, err := a.grpcClient.Login(rqt.Email, rqt.Password)
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, tokens)
}

func (a *AuthController) RefreshTokens(c *gin.Context) {
	token := c.Param("token")
	if len(token) == 0 {
		helper.BadRequestResponse(c, errors.New("todo"))

		return
	}

	tokens, err := a.grpcClient.RefreshTokens(token)
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, tokens)
}
