package redisrepository

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RateLimitRepository struct {
	client *redis.Client
}

func NewRateLimitRepository(client *redis.Client) *RateLimitRepository {
	return &RateLimitRepository{
		client: client,
	}
}

func (r *RateLimitRepository) Get(ctx context.Context, key string) (int64, error) {
	return r.client.Get(ctx, key).Int64()
}

func (r *RateLimitRepository) Increment(ctx context.Context, key string, ttl time.Duration) (int64, error) {
	pipe := r.client.Pipeline()
	count := pipe.Incr(ctx, key)
	pipe.ExpireNX(ctx, key, ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return 0, err
	}

	return count.Result()
}

func (r *RateLimitRepository) Reset(ctx context.Context, key string) error {
	return r.client.Del(ctx, key).Err()
}
