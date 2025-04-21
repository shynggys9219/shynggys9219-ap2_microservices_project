package app

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/usecase"
	"log"
	"os"
	"os/signal"
	"syscall"

	statisticsvc "github.com/shynggys9219/ap2-apis-gen-statistics-service/service/frontend/client_stats/v1"
	usersvc "github.com/shynggys9219/ap2-apis-gen-user-service/service/frontend/client/v1"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/config"
	grpcusersvcclient "github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/grpc/client"
	grpcstatisticssvcclient "github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/grpc/statistics"
	httpserver "github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/pkg/grpcconn"
)

const serviceName = "api-gateway"

type App struct {
	httpServer *httpserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	clientServiceGRPCConn, err := grpcconn.New(cfg.GRPC.GRPCClient.UserServiceURL)
	if err != nil {
		return nil, err
	}

	statsServiceGRPCConn, err := grpcconn.New(cfg.GRPC.GRPCClient.StatisticsServiceURL)
	if err != nil {
		return nil, err
	}

	clientServiceClient := grpcusersvcclient.NewClient(usersvc.NewClientServiceClient(clientServiceGRPCConn))
	clientStatisticClient := grpcstatisticssvcclient.NewClient(
		statisticsvc.NewClientStatisticsServiceClient(statsServiceGRPCConn),
	)

	clientUsecase := usecase.NewClient(clientServiceClient)
	clientStatisticUsecase := usecase.NewClientStatistic(clientStatisticClient)

	// http service
	httpServer := httpserver.New(cfg.Server, clientUsecase, clientStatisticUsecase)

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
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)

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
