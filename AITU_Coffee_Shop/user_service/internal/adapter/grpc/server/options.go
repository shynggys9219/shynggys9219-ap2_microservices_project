package server

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func (a *API) setOptions(ctx context.Context) []grpc.ServerOption {
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionAge:      a.cfg.MaxConnectionAge,
			MaxConnectionAgeGrace: a.cfg.MaxConnectionAgeGrace,
		}),
		grpc.UnaryInterceptor(AuthInterceptor("mySecretKey")),
		grpc.MaxRecvMsgSize(a.cfg.MaxRecvMsgSizeMiB * (1024 * 1024)), // MaxRecvSize * 1 MB
	}

	return opts
}
