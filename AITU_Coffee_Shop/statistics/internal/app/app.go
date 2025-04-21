package app

import (
	"context"
	"fmt"

	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shynggys9219/ap2_microservices_project/statistics/config"
	grpcserver "github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/grpc"
	mongorepo "github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/mongo"
	natshandler "github.com/shynggys9219/ap2_microservices_project/statistics/internal/adapter/nats/handler"
	"github.com/shynggys9219/ap2_microservices_project/statistics/internal/usecase"
	mongocon "github.com/shynggys9219/ap2_microservices_project/statistics/pkg/mongo"
	natsconn "github.com/shynggys9219/ap2_microservices_project/statistics/pkg/nats"
	natsconsumer "github.com/shynggys9219/ap2_microservices_project/statistics/pkg/nats/consumer"
)

const serviceName = "statistics-service"

type App struct {
	grpcServer         *grpcserver.API
	natsPubSubConsumer *natsconsumer.PubSub
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

	// Repository
	clientRepo := mongorepo.NewClient(mongoDB.Conn)

	clientUsecase := usecase.NewClient(clientRepo)

	// Nats consumers
	natsPubSubConsumer := natsconsumer.NewPubSub(natsClient)

	// nats handler
	clientHandler := natshandler.NewClient(clientUsecase)
	natsPubSubConsumer.Subscribe(natsconsumer.PubSubSubscriptionConfig{
		Subject: cfg.Nats.NatsSubjects.ClientEventSubject,
		Handler: clientHandler.Handler,
	})

	gRPCServer := grpcserver.New(
		cfg.Server.GRPCServer,
		clientUsecase,
	)

	app := &App{
		grpcServer:         gRPCServer,
		natsPubSubConsumer: natsPubSubConsumer,
	}

	return app, nil
}

func (a *App) Close(ctx context.Context) {
	err := a.grpcServer.Stop(ctx)
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
	a.grpcServer.Run(ctx, errCh)
	a.natsPubSubConsumer.Start(ctx, errCh)
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
