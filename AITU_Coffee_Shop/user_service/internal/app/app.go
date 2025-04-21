package app

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/nats/producer"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/config"
	grpcserver "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server"
	httpserver "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/http/service"
	mongorepo "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/usecase"
	mongocon "github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/mongo"

	natsconn "github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/nats"
)

const serviceName = "user-service"

type App struct {
	httpServer *httpserver.API
	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	// nats client
	log.Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	clientProducer := producer.NewClientProducer(natsClient, cfg.Nats.NatsSubjects.ClientEventSubject)

	// Repository
	aiRepo := mongorepo.NewAi(mongoDB.Conn)
	userRepo := mongorepo.NewClient(mongoDB.Conn)

	// UseCase
	userUsecase := usecase.NewUser(aiRepo, userRepo, clientProducer)

	// http service
	httpServer := httpserver.New(cfg.Server, userUsecase)

	gRPCServer := grpcserver.New(
		cfg.Server.GRPCServer,
		userUsecase,
	)

	app := &App{
		httpServer: httpServer,
		grpcServer: gRPCServer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}

	err = a.grpcServer.Stop(ctx)
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
	a.httpServer.Run(errCh)
	a.grpcServer.Run(ctx, errCh)

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
