package middleware

import (
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
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

		c.Abort()
	}
	userId, err := a.auth.VerifyTokenAndGetId(authHeader[1])
	if err != nil {
		helper.UnauthorizedResponse(c)

		c.Abort()
	}

	md := metadata.New(map[string]string{"uid": userId})

	c.Next()
}
