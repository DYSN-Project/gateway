package api

import (
	"context"
	"github.com/DYSN-Project/gateway/internal/helper"
	"github.com/DYSN-Project/gateway/internal/transport/api/endpoints/v1"
	"github.com/DYSN-Project/gateway/internal/transport/grpc"
	"github.com/DYSN-Project/gateway/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Server struct {
	httpSrv *http.Server
	logger  *logger.Logger
}

type RestServerInterface interface {
	StartServer()
	StopServer()
}

func NewServer(port string,
	areaGrpcClient *grpc.Area,
	authGrpcClient *grpc.Auth,
	logger *logger.Logger,
) *Server {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		helper.SuccessResponse(c, "pong")
	})

	v1.RegisterAuthEndpoints(router, authGrpcClient)
	v1.RegisterAreaEndpoints(router, areaGrpcClient)

	srv := &http.Server{
		Addr:     port,
		Handler:  router,
		ErrorLog: logger.ErrorLog,
	}

	return &Server{
		httpSrv: srv,
		logger:  logger,
	}
}

func (s *Server) StartServer() {
	if err := s.httpSrv.ListenAndServe(); err != nil {
		s.logger.ErrorLog.Fatalf("Failed to listen and serve: %+v", err)
	}
	s.logger.InfoLog.Println("Start REST server...")
}

func (s *Server) StopServer() {
	if err := s.httpSrv.Shutdown(context.Background()); err != nil {
		s.logger.ErrorLog.Fatalf("Failed stopped serve: %+v", err)
	}
	s.logger.InfoLog.Println("Shutting down REST server...")
}
