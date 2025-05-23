package main

import (
	"context"
	"log"

	"github.com/shynggys9219/ap2_microservices_project/emailer_service/config"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/internal/app"
)

func main() {
	ctx := context.Background()
	// TODO: add telemetry here when the topic of logging will be covered

	// Parse config
	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to parse config: %v", err)

		return
	}

	application, err := app.New(ctx, cfg)
	if err != nil {
		log.Println("failed to setup application:", err)

		return
	}

	err = application.Run()
	if err != nil {
		log.Println("failed to run application: ", err)

		return
	}
}
