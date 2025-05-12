package app

import (
	"context"
	"fmt"
	"github.com/mailersend/mailersend-go"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/adapter/mailer"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/shynggys9219/ap2_microservices_project/emailer_service/config"
	natshandler "github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/adapter/nats/handler"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/usecase"
	natsconn "github.com/shynggys9219/ap2_microservices_project/emailer_service/pkg/nats"
	natsconsumer "github.com/shynggys9219/ap2_microservices_project/emailer_service/pkg/nats/consumer"
)

const serviceName = "emailer-service"

type App struct {
	natsPubSubConsumer *natsconsumer.PubSub
}

func New(ctx context.Context, cfg *config.Config) (*App, error) {
	log.Println(fmt.Sprintf("starting %v service", serviceName))

	// nats client
	log.Println("connecting to NATS", "hosts", strings.Join(cfg.Nats.Hosts, ","))
	natsClient, err := natsconn.NewClient(ctx, cfg.Nats.Hosts, cfg.Nats.NKey, cfg.Nats.IsTest)
	if err != nil {
		return nil, fmt.Errorf("nats.NewClient: %w", err)
	}
	log.Println("NATS connection status is", natsClient.Conn.Status().String())

	// Nats consumers
	natsPubSubConsumer := natsconsumer.NewPubSub(natsClient)

	// mailersend clien
	mailerSend := mailer.NewMailer(mailersend.NewMailersend(cfg.MailerKey))

	// UseCase
	customerUsecase := usecase.NewCustomer(nil, mailerSend)
	// nats handler
	clientHandler := natshandler.NewCustomer(customerUsecase)
	natsPubSubConsumer.Subscribe(natsconsumer.PubSubSubscriptionConfig{
		Subject: cfg.Nats.NatsSubjects.CustomerEventSubject,
		Handler: clientHandler.Handler,
	})

	app := &App{
		natsPubSubConsumer: natsPubSubConsumer,
	}

	return app, nil
}

func (a *App) Close(_ context.Context) {
	a.natsPubSubConsumer.Stop()
}

func (a *App) Run() error {
	errCh := make(chan error, 1)
	ctx := context.Background()
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
		log.Println("graceful shutdown completed!")
	}

	return nil
}
