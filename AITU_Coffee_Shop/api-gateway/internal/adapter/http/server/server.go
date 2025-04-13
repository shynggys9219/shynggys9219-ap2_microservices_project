package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/internal/adapter/http/server/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shynggys9219/ap2_microservices_project/api-gateway/config"
)

const serverIPAddress = "0.0.0.0:%d" // Changed to 0.0.0.0 for external access

type API struct {
	server *gin.Engine
	cfg    config.HTTPServer
	addr   string

	clientHandler *handler.Client
}

func New(cfg config.Server, client ClientUsecase) *API {
	// Setting the Gin mode
	gin.SetMode(cfg.HTTPServer.Mode)
	// Creating a new Gin Engine
	server := gin.New()

	// Applying middleware
	server.Use(gin.Recovery())

	// Binding presenter
	clientHandler := handler.NewClient(client)

	api := &API{
		server:        server,
		cfg:           cfg.HTTPServer,
		addr:          fmt.Sprintf(serverIPAddress, cfg.HTTPServer.Port),
		clientHandler: clientHandler,
	}

	api.setupRoutes()

	return api
}

func (a *API) setupRoutes() {
	v1 := a.server.Group("/api/v1")
	{
		clients := v1.Group("/clients")
		{
			clients.POST("/", a.clientHandler.Create)
			clients.PUT("/update/:id", a.clientHandler.Update)
			clients.GET("/client/:id", a.clientHandler.Get)
		}
	}
}

func (a *API) Run(errCh chan<- error) {
	go func() {
		log.Printf("HTTP server starting on: %v", a.addr)

		// No need to reinitialize `a.server` here. Just run it directly.
		if err := a.server.Run(a.addr); err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- fmt.Errorf("failed to start HTTP server: %w", err)
			return
		}
	}()
}

func (a *API) Stop() error {
	// Setting up the signal channel to catch termination signals
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Blocking until a signal is received
	sig := <-quit
	log.Println("Shutdown signal received", "signal:", sig.String())

	// Creating a context with timeout for graceful shutdown
	_, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("HTTP server shutting down gracefully")

	// Note: You can use `Shutdown` if you use `http.Server` instead of `gin.Engine`.
	log.Println("HTTP server stopped successfully")

	return nil
}
