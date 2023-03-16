package grpc

import (
	"context"
	"github.com/DYSN-Project/gateway/internal/conversion"
	"github.com/DYSN-Project/gateway/internal/model/response"
	"github.com/DYSN-Project/gateway/internal/transport/grpc/pb/auth"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"google.golang.org/grpc"
)

type AuthClient interface {
	Register(email, password, lang string) (*response.User, error)
	ConfirmRegister(email, password, code string) (*response.Tokens, error)
	Login(email, password string) (*response.Tokens, error)
	UpdateTokens(token string) (*response.Tokens, error)
	GetUserByToken(token string) (*response.User, error)
	RecoveryPassword(email string) error
	RecoveryPasswordConfirm(email, code string) error
	ChangePassword(email, password string) error
}

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

func (a *Auth) Register(email, password, lang string) (*response.User, error) {
	rqt := &pb.RegisterRequest{
		Email:    email,
		Password: password,
		Lang:     lang,
	}

	user, err := a.client.Register(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	var mdl *response.User
	conversion.FromUserPbToUserModel(user, mdl)

	return mdl, nil
}

func (a *Auth) ConfirmRegister(email, password, code string) (*response.Tokens, error) {
	rqt := &pb.ConfirmRequest{
		Email:    email,
		Password: password,
		Code:     code,
	}

	tokens, err := a.client.ConfirmRegister(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	return response.NewTokens(tokens.GetAccessToken(), tokens.GetRefreshToken()), nil
}

func (a *Auth) Login(email, password string) (*response.Tokens, error) {
	rqt := &pb.LoginRequest{
		Email:    email,
		Password: password,
	}

	tokens, err := a.client.Login(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	return response.NewTokens(tokens.GetAccessToken(), tokens.GetRefreshToken()), nil
}

func (a *Auth) UpdateTokens(token string) (*response.Tokens, error) {
	rqt := &pb.Token{
		Token: token,
	}

	tokens, err := a.client.UpdateTokens(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	return response.NewTokens(tokens.GetAccessToken(), tokens.GetRefreshToken()), nil
}

func (a *Auth) GetUserByToken(token string) (*response.User, error) {
	rqt := &pb.Token{
		Token: token,
	}

	user, err := a.client.GetUserByToken(context.Background(), rqt)
	if err != nil {
		return nil, err
	}

	var mdl *response.User
	conversion.FromUserPbToUserModel(user, mdl)

	return mdl, nil
}

func (a *Auth) RecoveryPassword(email string) error {
	rqt := &pb.RemovePasswordRequest{
		Email: email,
	}
	_, err := a.client.RemovePassword(context.Background(), rqt)

	return err
}

func (a *Auth) RecoveryPasswordConfirm(email, code string) error {
	rqt := &pb.ConfirmRemovePasswordRequest{
		Email: email,
		Code:  code,
	}
	_, err := a.client.RemovePasswordConfirm(context.Background(), rqt)

	return err
}

func (a *Auth) ChangePassword(email, password string) error {
	rqt := &pb.ChangePasswordRequest{
		Email:    email,
		Password: password,
	}
	_, err := a.client.ChangePassword(context.Background(), rqt)

	return err
}
