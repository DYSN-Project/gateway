package grpc

import (
	"context"
	"github.com/DYSN-Project/gateway/internal/conversion"
	"github.com/DYSN-Project/gateway/internal/model/response"
	"github.com/DYSN-Project/gateway/internal/transport/grpc/pb/auth"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"google.golang.org/grpc"
)

type Auth struct {
	client pb.AuthClient
}

func NewAuth(port string, errLog *logger.Logger) *Auth {
	opt := []grpc.DialOption{
		grpc.WithInsecure(),
	}

	conn, err := grpc.Dial(port, opt...)
	if err != nil {
		errLog.ErrorLog.Fatal(err)
	}

	return &Auth{
		client: pb.NewAuthClient(conn),
	}
}

func (a *Auth) Register(email, password string) (string, error) {
	rqt := &pb.RegisterRequest{
		Email:    email,
		Password: password,
	}

	registerToken, err := a.client.Register(context.Background(), rqt)
	if err != nil {
		return "", err
	}

	return registerToken.GetToken(), nil
}

func (a *Auth) ConfirmRegister(token string) (*response.User, error) {
	rqt := &pb.Token{
		Token: token,
	}

	userRsp, err := a.client.ConfirmRegister(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	user := response.NewUserIngot()
	conversion.FromUserPbToUserModel(userRsp, user)

	return user, nil
}
