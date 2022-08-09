package middleware

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
	"strings"
)

const tokenKey = "Authorization"

type AuthMdl struct {
	auth *grpc.Auth
}

func NewAuthMdl(auth *grpc.Auth) *AuthMdl {
	return &AuthMdl{
		auth: auth,
	}
}

func (a *AuthMdl) Auth(c *gin.Context) {
	authHeader := strings.Split(c.GetHeader(tokenKey), "Bearer ")
	if len(authHeader) != 2 {
		helper.UnauthorizedResponse(c)
	}
	a.auth.ConfirmRegister(authHeader[1])

	c.Next()
}
