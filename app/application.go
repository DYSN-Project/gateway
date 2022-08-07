package app

import (
	"context"
	"github.com/DYSN-Project/gateway/config"
	"github.com/DYSN-Project/gateway/internal/transport/api"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) {
	cfg := config.NewConfig()
	lgr := logger.NewLogger()

	areaGrpcClient := grpc.NewArea(cfg.GetAreaPort(), lgr)
	authGrpcClient := grpc.NewAuth(cfg.GetAuthPort(), lgr)

	server := api.NewServer(cfg.GetSelfPort(), areaGrpcClient, authGrpcClient, lgr)

	go server.StartServer()
	defer server.StopServer()

	sgn := make(chan os.Signal, 1)
	signal.Notify(sgn, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-ctx.Done():
	case <-sgn:
	}
}
