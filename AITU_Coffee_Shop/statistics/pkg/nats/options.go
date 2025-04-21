package nats

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nkeys"
)

// Nats jetstream settings.
const (
	maxWaitFetch = 3 * time.Second
	maxPending   = 256
	maxReconnect = -1
)

func setOptions(ctx context.Context, hosts []string, nkey string, isTest bool) (nats.Options, error) {
	opts := nats.GetDefaultOptions()
	opts.Servers = hosts
	opts.Timeout = nats.DefaultTimeout
	opts.DrainTimeout = nats.DefaultDrainTimeout
	opts.PingInterval = nats.DefaultPingInterval
	opts.MaxPingsOut = nats.DefaultMaxPingOut
	opts.Verbose = true
	opts.RetryOnFailedConnect = true
	opts.MaxReconnect = maxReconnect
	opts.ReconnectWait = nats.DefaultReconnectWait

	opts.AsyncErrorCB = func(nc *nats.Conn, consumer *nats.Subscription, err error) {
		if err != nil {
			log.Println(
				ctx,
				"nats async error",
				"subject", consumer.Subject,
				"status", nc.Status().String(),
				err,
			)
		}
	}

	opts.ReconnectedCB = func(nc *nats.Conn) {
		log.Println(ctx, "nats reconnect status", "status", nc.Status().String())
	}

	opts.DisconnectedErrCB = func(nc *nats.Conn, err error) {
		log.Println(ctx, "nats disconnect status", "status", nc.Status().String())

		if err != nil {
			log.Println(ctx, "nats disconnect error", err.Error())
		}
	}

	opts.ClosedCB = func(nc *nats.Conn) {
		log.Println(ctx, "nats closed status",
			"status", nc.Status().String())
	}

	// if test run - skip authentication with nkey
	if isTest {
		return opts, nil
	}

	// auth with nkey
	kp, err := nkeys.FromSeed([]byte(nkey))
	if err != nil {
		return nats.Options{}, fmt.Errorf("failed to create KeyPair: %w", err)
	}

	publicKey, err := kp.PublicKey()
	if err != nil {
		return nats.Options{}, fmt.Errorf("failed to create public key: %w", err)
	}

	opts.Nkey = publicKey
	opts.SignatureCB = func(nonce []byte) ([]byte, error) {
		return kp.Sign(nonce)
	}

	return opts, nil
}
