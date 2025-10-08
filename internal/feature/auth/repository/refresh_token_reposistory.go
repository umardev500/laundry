package repository

import (
	"time"

	"github.com/umardev500/laundry/internal/app/appctx"
)

type RefreshTokenRepository interface {
	Set(ctx *appctx.Context, userID string, token string, ttl time.Duration) error
	Get(ctx *appctx.Context, userID string) (string, error)
	Delete(ctx *appctx.Context, userID string) error
}
