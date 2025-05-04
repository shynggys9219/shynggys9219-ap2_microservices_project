package app

import (
	"context"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/nats/producer"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/security"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/config"
	grpcserver "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/grpc/server"
	mongorepo "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/usecase"
	mongocon "github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/mongo"

	natsconn "github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/nats"
)

const serviceName = "user-service"

type App struct {
	grpcServer *grpcserver.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	// mongo transactor
	transactor := mongocon.NewTransactor(mongoDB.Client)

	// nats client
	log.Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	customerProducer := producer.NewCustomerProducer(natsClient, cfg.Nats.NatsSubjects.CustomerEventSubject)

	// Repository
	aiRepo := mongorepo.NewAi(mongoDB.Conn)
	customerRepo := mongorepo.NewCustomer(mongoDB.Conn)
	err = customerRepo.EnsureIndexes(ctx)
	if err != nil {
		log.Println("customerRepo.EnsureIndexes", err)
	}

	jwtManager := security.NewJWTManager(cfg.JWTManager.SecretKey)

	// UseCase
	customerUsecase := usecase.NewCustomer(aiRepo, customerRepo, customerProducer, transactor.WithinTransaction, jwtManager)

	// gRPC server
	gRPCServer := grpcserver.New(
		cfg.Server.GRPCServer,
		customerUsecase,
	)

	app := &App{
		grpcServer: gRPCServer,
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

	log.Println(fmt.Sprintf("service %v started", serviceName))

	// Waiting signal
	shutdownCh := make(chan os.Signal, 1)
	signal.Notify(shutdownCh, syscall.SIGINT, syscall.SIGTERM)

	select {
	case errRun := <-errCh:
		return errRun

	case s := <-shutdownCh:
		log.Println(fmt.Sprintf("received signal: %v. Running graceful shutdown...", s))
		log.Println("graceful shutdown completed!")
	}

	return nil
}
