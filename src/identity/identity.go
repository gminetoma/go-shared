package identity

import (
	"context"

	"github.com/gminetoma/go-shared/src/errs"
)

type contextKey string

const ownerIDKey contextKey = "ownerID"

func SetOwnerID(ctx context.Context, ownerID string) context.Context {
	return context.WithValue(ctx, ownerIDKey, ownerID)
}

func OwnerID(ctx context.Context) (string, bool) {
	ownerID, ok := ctx.Value(ownerIDKey).(string)
	return ownerID, ok
}

func Require(ctx context.Context) (string, error) {
	ownerID, ok := OwnerID(ctx)
	if !ok {
		return "", errs.ErrUnauthenticated
	}

	return ownerID, nil
}
