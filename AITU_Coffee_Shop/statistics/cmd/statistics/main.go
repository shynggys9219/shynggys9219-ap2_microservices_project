package main

import (
	"context"
	"github.com/shynggys9219/ap2-apis-gen-user-service/events/v1"
	"google.golang.org/protobuf/proto"
	"log"
	"strings"
	"time"

	"github.com/nats-io/nats.go"
)

type Client struct {
	ID              uint64
	Name            string
	Phone           string
	Email           string
	CurrentPassword string
	NewPassword     string
	PasswordHash    string
	NewPasswordHash string
	CreatedAt       time.Time
	UpdatedAt       time.Time

	IsDeleted bool
}

// Simulate proto.Unmarshal for your proto message
func UnmarshalCreatedEvent(data []byte) (*Client, error) {
	var msg events.Client
	err := proto.Unmarshal(data, &msg)
	if err != nil {
		return nil, err
	}

	return &Client{
		ID:        msg.Id,
		Name:      msg.Name,
		Phone:     msg.Phone,
		Email:     msg.Email,
		CreatedAt: msg.CreatedAt.AsTime(),
		IsDeleted: msg.IsDeleted,
	}, nil
}

func main() {
	natsHosts := "localhost:4222,localhost:4222,localhost:4222"
	subject := "ap2.user_svc.event.created"

	// Connect to NATS with default settings
	nc, err := nats.Connect(strings.Join(strings.Split(natsHosts, ","), ","))
	if err != nil {
		log.Fatalf("failed to connect to NATS: %v", err)
	}
	defer nc.Drain()

	log.Println("Connected to NATS")

	// Subscribe to the subject
	_, err = nc.Subscribe(subject, func(m *nats.Msg) {
		ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
		defer cancel()

		msg, err := UnmarshalCreatedEvent(m.Data)
		if err != nil {
			log.Printf("failed to parse message: %v", err)
			return
		}

		log.Printf("[Received] ID: %s | Name: %s", msg.ID, msg.Name)

		select {
		case <-ctx.Done():
			log.Println("context timeout while processing")
		default:
			log.Println("processed message")
		}
	})
	if err != nil {
		log.Fatalf("failed to subscribe to %s: %v", subject, err)
	}

	log.Printf("Listening on subject: %s", subject)
	select {}
}
