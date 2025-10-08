package repository

import (
	"time"

	"github.com/umardev500/laundry/internal/app/appctx"
	"github.com/umardev500/laundry/internal/infra/database/redis"
)

type redisRefreshToken struct {
	client *redis.RedisClient
}

func NewRedisRefreshTokenRepository(client *redis.RedisClient) RefreshTokenRepository {
	return &redisRefreshToken{
		client: client,
	}
}

// Delete implements RefreshTokenRepository.
func (r *redisRefreshToken) Delete(ctx *appctx.Context, userID string) error {
	key := newRefreshTokenKey(userID)
	return r.client.Del(ctx, key.String()).Err()
}

// Get implements RefreshTokenRepository.
func (r *redisRefreshToken) Get(ctx *appctx.Context, userID string) (string, error) {
	key := newRefreshTokenKey(userID)
	return r.client.Get(ctx, key.String()).Result()
}

// Set implements RefreshTokenRepository.
func (r *redisRefreshToken) Set(ctx *appctx.Context, userID string, token string, ttl time.Duration) error {
	key := newRefreshTokenKey(userID)
	return r.client.Set(ctx, key.String(), token, ttl).Err()
}
