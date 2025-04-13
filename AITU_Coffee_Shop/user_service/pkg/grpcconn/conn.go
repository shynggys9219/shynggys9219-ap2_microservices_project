package grpcconn

import (
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	MaxRecvMsgSize = 24 * 1024 * 1024 // 24 MB
	timeKeepalive  = 10 * time.Second
)

// New initializes a new gRPC client connection without interceptors.
func New(target string) (*grpc.ClientConn, error) {
	var keepaliveClientParameters = keepalive.ClientParameters{
		Time:                timeKeepalive, // Send pings every 10 seconds if there is no activity
		Timeout:             time.Second,   // Wait 1 second for ping ack before considering the connection dead
		PermitWithoutStream: true,          // Send pings even without active streams
	}

	// Create the gRPC client connection with necessary options
	clientConn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),             // Insecure connection, no TLS
		grpc.WithKeepaliveParams(keepaliveClientParameters),                  // Keepalive configuration
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxRecvMsgSize)), // Set max message size
	)
	if err != nil {
		return nil, fmt.Errorf("grpc client: %w", err)
	}

	return clientConn, nil
}
