package consumer

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/nats-io/nats.go"
	natscl "github.com/shynggys9219/ap2_microservices_project/emailer_service/pkg/nats"
	"github.com/shynggys9219/ap2_microservices_project/emailer_service/pkg/safe"
)

type PubSub struct {
	subsCfg []PubSubSubscriptionConfig
	client  *natscl.Client
	subs    []*nats.Subscription
	wg      *sync.WaitGroup
	stop    chan struct{}
}

type PubSubSubscriptionConfig struct {
	Subject string
	Handler natscl.MsgHandler
}

func NewPubSub(client *natscl.Client) *PubSub {
	return &PubSub{
		client: client,
		subs:   make([]*nats.Subscription, 0),
		wg:     &sync.WaitGroup{},
		stop:   make(chan struct{}),
	}
}

func (c *PubSub) Subscribe(cfg ...PubSubSubscriptionConfig) {
	c.subsCfg = append(c.subsCfg, cfg...)
}

// Start is used to start pulling messages from all subjects.
func (c *PubSub) Start(ctx context.Context, errCh chan<- error) {
	for i := range c.subsCfg {
		c.wg.Add(1)
		go safe.Do(ctx, func() {
			defer c.wg.Done()
			err := c.consume(c.subsCfg[i])
			if err != nil {
				errCh <- fmt.Errorf("failed to start consuming NATS subject %v: %w", c.subsCfg[i].Subject, err)
			}
		})
	}
}

func (c *PubSub) Stop() {
	c.wg.Wait()

	for _, sub := range c.subs {
		err := sub.Unsubscribe()
		if err != nil {
			log.Println("failed to unsubscribe",
				"subject:", sub.Subject,
				"error:", err,
			)
		}
	}
}

func (c *PubSub) consume(cfg PubSubSubscriptionConfig) error {
	sub, err := c.client.Subscribe(cfg.Subject, cfg.Handler)
	if err != nil {
		return fmt.Errorf("c.client.Subscribe: %w", err)
	}

	c.subs = append(c.subs, sub)

	log.Println("consuming NATS subject started", "subject: ", cfg.Subject)

	return nil
}
