package grpc

import (
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
