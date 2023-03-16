package v1

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/model/form"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
)

type AuthInterface interface {
	Register(c *gin.Context)
	ConfirmRegister(c *gin.Context)
	Login(c *gin.Context)
	RefreshTokens(c *gin.Context)
	RecoveryPassword(c *gin.Context)
	RecoveryPasswordConfirm(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type AuthController struct {
	grpcClient *grpc.Auth
}

func NewAuthController(grpcClient *grpc.Auth) *AuthController {
	return &AuthController{
		grpcClient: grpcClient,
	}
}

func (a *AuthController) Register(c *gin.Context) {
	rqt := &form.Register{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	user, err := a.grpcClient.Register(rqt.Email,
		rqt.Password,
		rqt.Lang)

	if err != nil {
		helper.ErrorResponse(c, err)
	}

	helper.SuccessResponse(c, user)
}

func (a *AuthController) ConfirmRegister(c *gin.Context) {
	rqt := &form.ConfirmRegister{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	tokens, err := a.grpcClient.ConfirmRegister(rqt.Email,
		rqt.Password,
		rqt.Code)
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, tokens)
}

func (a *AuthController) Login(c *gin.Context) {
	rqt := &form.Login{}
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
	rqt := &form.Token{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	tokens, err := a.grpcClient.UpdateTokens(rqt.Token)
	if err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, tokens)
}

func (a *AuthController) RecoveryPassword(c *gin.Context) {
	rqt := &form.RecoveryPasswordRqt{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	if err := a.grpcClient.RecoveryPassword(rqt.Email); err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, "success remove request.Check your email")
}

func (a *AuthController) RecoveryPasswordConfirm(c *gin.Context) {
	rqt := &form.RecoveryPasswordConfirm{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	if err := a.grpcClient.RecoveryPasswordConfirm(rqt.Email, rqt.Code); err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, "recovery password confirmed")
}

func (a *AuthController) ChangePassword(c *gin.Context) {
	rqt := &form.ChangePassword{}
	if err := c.BindJSON(rqt); err != nil {
		helper.BadRequestResponse(c, err)

		return
	}

	if err := a.grpcClient.ChangePassword(rqt.Email, rqt.Password); err != nil {
		helper.ErrorResponse(c, err)

		return
	}

	helper.SuccessResponse(c, "password changed")
}
