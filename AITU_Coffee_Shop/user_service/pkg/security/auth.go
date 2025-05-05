package security

import (
	"context"
	"google.golang.org/grpc/metadata"
	"strings"
)

func TokenFromCtx(ctx context.Context) (string, bool) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", false
	}

	authHeader := md["authorization"]
	if len(authHeader) == 0 {
		return "", false
	}

	tokenStr := strings.TrimPrefix(authHeader[0], "Bearer ")

	return tokenStr, true
}
