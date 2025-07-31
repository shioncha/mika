package repository

import (
	"context"
	"time"
)

type Session struct {
	ID         string
	UserID     string
	DeviceInfo string
	IPAddress  string
	CreatedAt  time.Time
	LastUsedAt time.Time
}

type SessionRepository interface {
	// 新規セッションを作成
	CreateSession(ctx context.Context, userID, deviceInfo, ipAddress string, ttl time.Duration) (refreshToken string, err error)

	// セッションを取得
	GetSession(ctx context.Context, refreshToken string) (*Session, error)

	// セッションを更新
	ExtendSession(ctx context.Context, refreshToken string, ttl time.Duration) error

	// セッションを削除
	DeleteSession(ctx context.Context, refreshToken string) error

	// ユーザーのセッション一覧を取得
	GetUserSessions(ctx context.Context, userID string) ([]*Session, error)
}
