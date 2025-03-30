package app

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/shynggys9219/ap2_microservices_project/user_svc/config"
	httpservice "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/http/service"
	mongorepo "github.com/shynggys9219/ap2_microservices_project/user_svc/internal/adapter/mongo"
	"github.com/shynggys9219/ap2_microservices_project/user_svc/internal/usecase"
	mongocon "github.com/shynggys9219/ap2_microservices_project/user_svc/pkg/mongo"
)

const serviceName = "user-service"

type App struct {
	httpServer *httpservice.API
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	log.Println("connecting to mongo", "database", cfg.Mongo.Database)
	mongoDB, err := mongocon.NewDB(ctx, cfg.Mongo)
	if err != nil {
		return nil, fmt.Errorf("mongo: %w", err)
	}

	// Repository
	aiRepo := mongorepo.NewAi(mongoDB.Conn)
	userRepo := mongorepo.NewClient(mongoDB.Conn)

	// UseCase
	userUsecase := usecase.NewUser(aiRepo, userRepo)

	// http service
	httpServer := httpservice.New(cfg.Server, userUsecase)

	app := &App{
		httpServer: httpServer,
	}

	return app, nil
}

func (a *App) Close() {
	err := a.httpServer.Stop()
	if err != nil {
		log.Println("failed to shutdown gRPC service", err)
	}
}

func (a *App) Run() error {
	errCh := make(chan error, 1)

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

		a.Close()
		log.Println("graceful shutdown completed!")
	}

	return nil
}
