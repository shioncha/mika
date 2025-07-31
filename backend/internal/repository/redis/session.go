package redisrepository

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
	"github.com/shioncha/mika/backend/internal/repository"
)

type SessionRepository struct {
	client *redis.Client
}

func NewSessionRepository(client *redis.Client) *SessionRepository {
	return &SessionRepository{
		client: client,
	}
}

func (r *SessionRepository) CreateSession(ctx context.Context, userID, deviceInfo, ipAddress string, ttl time.Duration) (string, error) {
	rb := make([]byte, 32)
	if _, err := rand.Read(rb); err != nil {
		return "", err
	}
	refreshToken := base64.URLEncoding.EncodeToString(rb)

	hash := sha256.Sum256([]byte(refreshToken))
	hashStr := fmt.Sprintf("%x", hash)
	sessionKey := "session:" + hashStr

	pipe := r.client.Pipeline()
	pipe.HSet(ctx, sessionKey, map[string]interface{}{
		"user_id":      userID,
		"device_info":  deviceInfo,
		"ip_address":   ipAddress,
		"last_used_at": time.Now().Format(time.RFC3339),
	})
	pipe.Expire(ctx, sessionKey, ttl)

	userSessionsKey := "user_sessions:" + userID
	pipe.SAdd(ctx, userSessionsKey, hashStr)

	if _, err := pipe.Exec(ctx); err != nil {
		return "", err
	}

	return refreshToken, nil
}

func (r *SessionRepository) GetSession(ctx context.Context, refreshToken string) (*repository.Session, error) {
	hash := sha256.Sum256([]byte(refreshToken))
	hashStr := fmt.Sprintf("%x", hash)
	sessionKey := "session:" + hashStr

	data, err := r.client.HGetAll(ctx, sessionKey).Result()
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return nil, fmt.Errorf("session not found")
	}

	createdAt, _ := time.Parse(time.RFC3339Nano, data["created_at"])
	lastUsedAt, _ := time.Parse(time.RFC3339Nano, data["last_used_at"])

	return &repository.Session{
		ID:         hashStr,
		UserID:     data["user_id"],
		DeviceInfo: data["device_info"],
		IPAddress:  data["ip_address"],
		CreatedAt:  createdAt,
		LastUsedAt: lastUsedAt,
	}, nil
}

func (r *SessionRepository) ExtendSession(ctx context.Context, refreshToken string, ttl time.Duration) error {
	hash := sha256.Sum256([]byte(refreshToken))
	hashStr := fmt.Sprintf("%x", hash)
	sessionKey := "session:" + hashStr

	pipe := r.client.Pipeline()
	pipe.HSet(ctx, sessionKey, "last_used_at", time.Now().Format(time.RFC3339))
	pipe.Expire(ctx, sessionKey, ttl)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) DeleteSession(ctx context.Context, refreshToken string) error {
	session, err := r.GetSession(ctx, refreshToken)
	if err != nil {
		if err.Error() == "session not found" {
			return nil
		}
		return err
	}

	hash := sha256.Sum256([]byte(refreshToken))
	hashStr := fmt.Sprintf("%x", hash)
	sessionKey := "session:" + hashStr

	pipe := r.client.Pipeline()
	pipe.Del(ctx, sessionKey)

	userSessionsKey := "user_sessions:" + session.UserID
	pipe.SRem(ctx, userSessionsKey, hashStr)

	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	return nil
}

func (r *SessionRepository) GetUserSessions(ctx context.Context, userID string) ([]*repository.Session, error) {
	userSessionsKey := "user_sessions:" + userID
	sessionIDs, err := r.client.SMembers(ctx, userSessionsKey).Result()
	if err != nil {
		return nil, err
	}

	var sessions []*repository.Session
	for _, sessionID := range sessionIDs {
		session, err := r.GetSession(ctx, sessionID)
		if err != nil {
			continue // Ignore errors for individual sessions
		}
		sessions = append(sessions, session)
	}

	if len(sessions) == 0 {
		return nil, fmt.Errorf("no sessions found for user %s", userID)
	}

	return sessions, nil
}
