package transactor

import "context"

type WithinTransactionFunc func(ctx context.Context, fn func(fnCtx context.Context) error) error
