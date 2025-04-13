package app

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"

	svc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/config"
	grpcclient "github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/grpc/client"
	httpserver "github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/pkg/grpcconn"
)

const serviceName = "user-service"

type App struct {
	httpServer *httpserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	clientServiceGRPCConn, err := grpcconn.New(cfg.GRPC.GRPCClient.UserServiceURL)
	if err != nil {
		return nil, err
	}

	clientServiceClient := grpcclient.NewClient(svc.NewClientServiceClient(clientServiceGRPCConn))

	clientUsecase := usecase.NewClient(clientServiceClient)

	// http service
	httpServer := httpserver.New(cfg.Server, clientUsecase)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
	a.httpServer.Run(errCh)

	log.Println(fmt.Sprintf("service %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Println(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))

		a.Close(ctx)
		log.Println("graceful shutdown completed!")
	}

	return nil
}
