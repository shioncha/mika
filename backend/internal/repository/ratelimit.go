package repository

import (
	"context"
	"time"
)

type RateLimitRepository interface {
	// Keyのカウンタを取得
	Get(ctx context.Context, key string) (int64, error)

	// Keyのカウンターを増やす
	Increment(ctx context.Context, key string, ttl time.Duration) (int64, error)

	// Keyのカウンターをリセット
	Reset(ctx context.Context, key string) error
}
