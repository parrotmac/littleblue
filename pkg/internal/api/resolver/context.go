package resolver

import (
	"context"

	"github.com/pkg/errors"

	"github.com/parrotmac/littleblue/pkg/types"
)

func getUserIDFromContext(ctx context.Context) (string, error) {
	userIDIface := ctx.Value(types.UserIDCtxKey)
	userID, ok := userIDIface.(string)
	if !ok {
		return "", errors.New(types.ErrNotFound)
	}
	return userID, nil
}
