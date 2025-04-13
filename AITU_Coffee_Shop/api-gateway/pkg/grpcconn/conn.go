package grpcconn

import (
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

const (
	MaxRecvMsgSize = 12 * 1024 * 1024
	timeKeepalive  = 10 * time.Second
)

func New(target string) (*grpc.ClientConn, error) {
	var keepaliveClientParameters = keepalive.ClientParameters{
		Time:                timeKeepalive,
		Timeout:             time.Second,
		PermitWithoutStream: true,
	}

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithKeepaliveParams(keepaliveClientParameters),
		grpc.WithDefaultCallOptions(grpc.MaxCallRecvMsgSize(MaxRecvMsgSize)),
	}

	return grpc.Dial(target, opts...)
}
